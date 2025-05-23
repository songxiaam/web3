package ws

//🔑 核心收发流程
//客户端连接到 WebSocket → Server.ReadAndWrite() 注册并启动读写心跳流程
//写协程 持续监听 Server.Send 通道，将外部推送的数据通过 WebSocket 发给客户端
//读协程 监听客户端消息，仅处理 “ping” 心跳
//心跳检查：无 ping 则超时断开
//全局广播：StartServer 抓取币价更新，广播给所有 Server.Send
//这样就实现了一个可靠的心跳保活 + 推送最新行情的 WebSocket 服务架构。

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"pledgev2-backend/config"
	"pledgev2-backend/log"
	"pledgev2-backend/pkg/models/kucoin"
	"sync"
	"time"
)

const SuccessCode = 0
const PongCode = 1
const ErrorCode = -1

type Server struct {
	sync.Mutex
	Id       string
	Socket   *websocket.Conn
	Send     chan []byte
	LastTime int64
}

type ServerManager struct {
	Servers    sync.Map
	Broadcast  chan []byte
	Register   chan *Server
	Unregister chan *Server
}

type Message struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

var Manager = ServerManager{}
var UserPingPongDurTime = config.Config.Env.WssTimeoutDuration

func (s *Server) SendToClient(data string, code int) {
	s.Lock()
	defer s.Unlock()
	// struct转json
	dataBytes, err := json.Marshal(Message{Code: code, Data: data})
	if err != nil {
		log.Logger.Error(err.Error())
	}
	// 发送消息
	err = s.Socket.WriteMessage(websocket.TextMessage, dataBytes)
	if err != nil {
		log.Logger.Sugar().Error(err.Error())
	}
}

// ReadAndWrite 维持 WebSocket 连接的“心跳”，防止客户端或网络长时间无数据导致连接被中间代理（如 nginx）或浏览器断掉
func (s *Server) ReadAndWrite() {
	errChan := make(chan error)
	Manager.Servers.Store(s.Id, s)

	defer func() {
		Manager.Servers.Delete(s) //从Servers map中移除该连接
		_ = s.Socket.Close()      // 关闭WebSocket
		close(s.Send)             // 关闭Send通道
	}()

	// 写协程(协程: 请量级线程)
	go func() {
		for {
			select {
			case message, ok := <-s.Send: // 不断从s.Send通道中取出待推送消息
				if !ok {
					errChan <- errors.New("Write message error") // 向 errChan 发送错误, 退出写协程
					return
				}
				// 正常情况下执行, 消息推送到客户端
				s.SendToClient(string(message), SuccessCode)
			}
		}
	}()

	// 读协程
	go func() {
		for {
			// 接收客户端消息
			_, message, err := s.Socket.ReadMessage()
			if err != nil {
				log.Logger.Sugar().Error(s.Id+" ReadMessage error", err)
				errChan <- err
				return
			}

			// 收到 ping 心跳探测包, 由客户端定时发出
			// 返回 pong 心跳响应包, 由服务端收到ping后马上返回给客户端
			// 非业务消息,用于监听连接是否中断
			if string(message) == "ping" || string(message) == `"ping"` || string(message) == `'ping''` {
				// 更新最后活跃时间
				s.LastTime = time.Now().Unix()
				// 回复pong
				s.SendToClient("pong", PongCode)
			}
			continue
		}
	}()

	// 主协程, 死循环
	for {
		select {
		case <-time.After(time.Second):
			// 每秒检查一次
			if time.Now().Unix()-s.LastTime >= UserPingPongDurTime {
				// 长时间不心跳, 报错
				s.SendToClient("heartbeat timeout", ErrorCode)
				return
			}
		case err := <-errChan:
			// 接收到错误
			log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
			return
		}
	}
}

func StartServer() {
	log.Logger.Info("WSServer start")
	for {
		select {
		case price, ok := <-kucoin.PlgrPriceChan:
			// 监听kucoin.PlgrPriceChan
			if ok {
				// 遍历所有活跃连接,将最新价格推送给客户端
				Manager.Servers.Range(func(key, value interface{}) bool {
					value.(*Server).SendToClient(price, SuccessCode)
					return true
				})
			}
		}
	}
}
