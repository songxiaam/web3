package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin-study/routers"
	"gin-study/test"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	gin.ForceConsoleColor()
	//gin.Default()
	//fmt.Println("start....")
	//database.InitDB()
	//
	//routers := routers.SetupRouter()
	//routers.Run(":8080")
	//fmt.Println("服务器启动成功：http://localhost:8080/user/login")

	//testBindCustomStruct()
	//testBindSrarchForm()
	//testBindUri()
	//testCookie()
	//testLog()
	//testMiddleWare()
	//testValidator()
	//testRouterLogFormat()
	//testGoroutine()
	//testRestart()
	//testPush()
	//testRouterParam()
	//testPureJSON()
	//testRedirect()
	//testRender()
	//testMultiServer()
	//testSecureJSON()
	//testHeader()
	//testReader()
	//testLetsEncrypt()
	//testUpload()
	//testSecret()
	//testLogFile()
	//testNil()
	//testChan()
	//testGoroutine2()
	//testContext()
	//testType()
	//testSelect()
	//testPanic()
	//testNew()
	//testSlice()
	//testPointer()
	//testA()
	//testMutex()
	//testRWMutex()
	//testWaitGroup()
	//testSyncMap()
	//testVar()
	//testChannel()
	//testChannelTimeout()
	//testString()
	//testJson()
	//testError()
	//testSyncPool()
	//testChannel2()
	//testChannel3()
	testDtm()
}

func testDtm() {
	go testRouter()
	time.Sleep(1 * time.Second)
	//dtm_demo.SubmitOrder()
	select {}
}

func testChannel3() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			return
		}()
	}
	time.Sleep(3 * time.Second)

	count := runtime.NumGoroutine()
	fmt.Println(count)
}

func testChannel2() {
	countChan := make(chan int) // 通道用于传递“加法请求”
	done := make(chan struct{}) // 通道用于通知“加完了”
	var result int

	// 专门的 goroutine 负责对 result 累加（其他人不能直接改 result）
	go func() {
		for val := range countChan {
			fmt.Println("1--->", val)
			result += val
			fmt.Println("2")
		}
		done <- struct{}{} // 通知 main：任务完成
	}()

	// 启动多个 goroutine 发起加法请求
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("-0--->", i)
			countChan <- 1 // 所有加法通过通道传递，避免竞争
			fmt.Println("0--->", i)
		}(i)
	}

	time.Sleep(100 * time.Millisecond) // 等待写入完成
	close(countChan)                   // 关闭通道，表示不再写入
	<-done                             // 等待加法协程完成
	fmt.Println("Final counter:", result)
}

func testSyncPool() {
	pool := sync.Pool{
		New: func() interface{} {
			return new(string)
		},
	}
	s1 := pool.Get().(*string)
	*s1 = "99"
	fmt.Println(*s1)

	pool.Put(s1)
	s2 := pool.Get().(*string)
	*s2 = "98"
	fmt.Println(*s2)
	//time.Sleep(1 * time.Second)
	pool.Put(s1)
	pool.Put(s2)
	s3 := pool.Get().(*string)
	fmt.Println(*s3)
	s4 := pool.Get().(*string)
	fmt.Println(*s4)
	s5 := pool.Get().(*string)
	fmt.Println(*s5)
}

func testError() {
	err := errors.New("error message")
	is := errors.Is(err, errors.New("error message"))
	fmt.Println(is)
	var notFound = errors.New("not found")

	notFoundErr := fmt.Errorf("err:%w", errors.New("not found"))
	is = errors.As(notFoundErr, &notFound)
	fmt.Println(is)
}

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func testJson() {
	jsonStr := `{"name":"Alice","age":30,"email":"alice@example.com"}`
	fmt.Println([]byte(jsonStr))
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name)

	marshal, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
}

func testString() {

	s := "hello world dsop sdp sdop"

	n := strings.SplitN(s, "o", 5)
	for _, v := range n {
		fmt.Println(v)
	}
}

func testChannelTimeout() {
	var ch = make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
	LOOP:
		for {
			select {
			case a, ok := <-ch:
				fmt.Println("接收--->", a, ok)
			case <-time.After(2 * time.Second):
				fmt.Println("timeout")
				break LOOP
			}
		}
		fmt.Println("fun1 over")
	}()

	go func() {
		defer wg.Done()
		//defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Println("发送--->", i)
			ch <- i
			fmt.Println("发送成功--->", i)
		}
	}()
	wg.Wait()
	fmt.Println("end")
}

func testChannel() {
	var ch = make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("接收--->")
			time.Sleep(100 * time.Millisecond)
			a, ok := <-ch
			fmt.Println("接收--->", a, ok)
			if !ok {
				break
			}
		}

	}()

	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Println("发送--->", i)
			ch <- i
			fmt.Println("发送成功--->", i)
		}
	}()
	wg.Wait()
	fmt.Println("end")
}

func testVar() {
	var i int
	fmt.Println(i)
	fmt.Println(&i)

	// new return pointer
	var j = new(int)
	fmt.Println(j)
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(&j)

	var c = make(chan int, 2)
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
	c <- 1
	a, ok := <-c
	fmt.Println(a, ok)
	close(c)
	a, ok = <-c

	fmt.Println(a, ok)
}

func testSyncMap() {

	var m map[any]int
	m["a"] = 1
	m[2] = 3
	fmt.Println(m)

	var sm sync.Map

	var sw sync.WaitGroup
	//var m map[int]int
	for i := 0; i < 100; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()

			sm.LoadOrStore(i, i*i)
			//m[i] = i * 1
		}()
	}
	sw.Wait()
	sm.LoadAndDelete(1)
	//fmt.Println(sm.Range)
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func testWaitGroup() {
	var wg sync.WaitGroup
	//wg.Add(10)
	fmt.Println("testWaitGroup start", time.Now().Format(time.DateTime))
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("--->", i, "func1 get lock at ", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(500 * time.Millisecond)
			fmt.Println("--->", i, "func1 release lock at ", time.Now().Format("2006-01-02 15:04:05"))
		}(i)
	}
	wg.Wait()
	fmt.Println("testWaitGroup end", time.Now().Format(time.DateTime))
}

func testRWMutex() {
	var lock sync.RWMutex
	fmt.Println("testMutex start")
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer lock.RUnlock()
			lock.RLock()
			fmt.Println("--->", i, "func1 get lock at ", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(500 * time.Millisecond)
			fmt.Println("--->", i, "func1 release lock at ", time.Now().Format("2006-01-02 15:04:05"))
		}(i)
	}

	time.Sleep(10 * time.Second)
}

func testMutex() {
	var lock sync.Mutex
	defer fmt.Println("testMutex over")
	go func() {
		defer func() {
			lock.Unlock()
			fmt.Println("func1  unlock at ", time.Now().Format("2006-01-02 15:04:05"))
		}()
		lock.Lock()
		fmt.Println("func1 get lock at ", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(1 * time.Second)
		fmt.Println("func1 release lock at ", time.Now().Format("2006-01-02 15:04:05"))
	}()
	time.Sleep(100 * time.Millisecond)
	go func() {
		defer func() {
			lock.Unlock()
			fmt.Println("func2  unlock at ", time.Now().Format("2006-01-02 15:04:05"))
		}()
		lock.Lock()
		fmt.Println("func2 get lock at ", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(1 * time.Second)
		fmt.Println("func2 release lock at ", time.Now().Format("2006-01-02 15:04:05"))
	}()
	time.Sleep(3.0 * time.Second)
	fmt.Println("over")
}

func testA() {
	a := 1
	var ai atomic.Uint32
	ai.Store(32)
	fmt.Println(ai.Load())

	ai.Swap(23)
	fmt.Println("-----")
	fmt.Println(ai.Load())
	old := ai.Add(^uint32(12) + 1)
	fmt.Println(old)
	fmt.Println(ai.Load())
	fmt.Println("-----")

	ok := ai.CompareAndSwap(23, 1000)
	fmt.Println(ok)
	fmt.Println(ai.Load())

	for i := 0; i < 10; i++ {
		go func() {

			b := a
			time.Sleep(1 * time.Second)
			a = b + 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(a)

	var b int32
	atomic.StoreInt32(&b, 123)
	fmt.Println(b)

}

var once sync.Once

func testOnce() {
	once.Do(func() {

	})
}

func testSyncMutex() {
	//var mu sync.Mutex
	//var wg sync.WaitGroup
	//var m sync.Map
}

func testOS() {

}

func testPointer() {
	var p *int
	a := 1
	p = &a
	fmt.Println(*p)

	var p2 unsafe.Pointer
	p2 = unsafe.Pointer(&a)
	fmt.Println(*(*int)(p2))

}

func testSlice() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[2:3]
	fmt.Println(s)
	s1 := append(s, 6)
	fmt.Println(s1)
	fmt.Println(arr)
}

func testNew() {
	ch := new(chan int)
	fmt.Println(ch)

	i := new(int)
	fmt.Println(i)
	fmt.Println(*i)

}

func testPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	willPanic()

}

func willPanic() {
	defer func() {
		fmt.Println("willPanic painc")
		if r := recover(); r != nil {
			fmt.Println("Recovered in willPanic ", r)
		}
	}()
	fmt.Println("willPanic")
	panic("出错啦啦")
	fmt.Println("willPanic2")
}

func testSelect() {

	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	fmt.Println("start")
	go func() {
		for {
			//time.Sleep(time.Second * 2)
			select {
			case val := <-jobs:
				fmt.Println("case1 ", val)
			case results <- 1:
				fmt.Println("case2")
			case <-ctx.Done():
				fmt.Println("testSelect Done")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("case3 return")
				//return
				//default:
				//	fmt.Println("default")
				//	time.Sleep(time.Second * 1)
			}
		}
	}()
	cancel()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		jobs <- i
	}
	time.Sleep(2 * time.Second)
}

type stypeA string
type stypeB string

func testType() {
	var k1 stypeA = "123"
	var k2 stypeB = "123"
	fmt.Println(k1, k2)
	fmt.Println(reflect.TypeOf(k1), reflect.TypeOf(k2))
}

func testContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	context.WithCancel(context.Background())
	ctx1 := context.WithValue(ctx, "key", "value")
	ctx2 := context.WithValue(ctx1, "key", "value2")
	ctx3 := context.WithValue(ctx2, "key", "value3")
	ctx4 := context.WithValue(ctx3, "key", "value4")

	v := ctx4.Value("key")
	fmt.Println(v)
	v = ctx3.Value("key")
	fmt.Println(v)
	v = ctx2.Value("key")
	fmt.Println(v)
	v = ctx1.Value("key")
	fmt.Println(v)

	//defer cancel()
	//done := ctx.Done()
	//fmt.Println(done)

	s := make(chan struct{})

	go func() {
		s <- struct{}{}
	}()
	s1 := <-s
	fmt.Println(s1)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done(): // 监听 context 的 Done 通道
			fmt.Println("Goroutine cancelled:", ctx.Err())
		}
	}(ctx)

	cancel()
	time.Sleep(time.Second)

}

func testGoroutine2() {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int)
	go worker2(ctx, jobs)

	jobs <- 1
	jobs <- 2
	jobs <- 3

	cancel()
	close(jobs)

	time.Sleep(1 * time.Second)
}

func worker2(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("jobs channel closed, exiting")
				return
			}
			fmt.Println("Processing job", job)
		case <-ctx.Done():
			fmt.Println("Received cancel signal, exiting")
			return
		}
	}
}

func testChan2() {
	jobs := make(chan int)    // 无缓冲通道
	results := make(chan int) // 无缓冲通道

	// 启动三个 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	fmt.Println("before send")

	// 向 jobs 通道发送任务
	go func() {
		for j := 1; j <= 5; j++ {
			fmt.Printf("Sending job %d\n", j)
			jobs <- j // 发送时会阻塞，直到有 worker 接收
		}
		close(jobs) // 关闭任务通道，告知 worker 没有新任务了
		fmt.Println("close jobs")
	}()

	// 收集结果
	for a := 1; a <= 5; a++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}

func testChan() {
	jobs := make(chan int, 4)
	results := make(chan int, 4)

	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}
	fmt.Println("before send")
	for j := 1; j <= 11; j++ {
		jobs <- j
		//fmt.Println(<-results)
	}
	fmt.Println("before close")
	close(jobs)
	v, ok := <-jobs
	fmt.Println(v, ok)
	fmt.Println("close")
	time.Sleep(time.Millisecond * 10000)
	//for a := 1; a <= 2; a++ {
	//	fmt.Println(<-results)
	//}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Printf("----> Worker %d started\n", id)
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2
		fmt.Printf("Worker %d send %d\n", id, job)
	}
}

func testNil() {
	//var a, b interface{} = nil, nil
	//fmt.Println(a == b)
	//
	//var c *[]int = nil
	//var d interface{} = nil
	//fmt.Println(c == d)
	//
	//const name = "123"

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k := range m {
		fmt.Println(k)
		if k == "c" {
			delete(m, k)
			m["d"] = 998
		}
	}
	fmt.Println(m)

	ch := make(chan int)

	go func() {
		ch <- 2
	}()

	go func() {
		ch <- 1
	}()

	//val := <-ch
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func testBindCustomStruct() {
	router := gin.Default()
	router.GET("/getb", test.GetDataB)
	router.GET("/getc", test.GetDataC)
	router.GET("/getd", test.GetDataD)

	router.Run()
}

func testBindSrarchForm() {
	router := gin.Default()
	router.GET("/testing", test.StartPage)
	router.Run(":8085")
}

func testBindUri() {
	router := gin.Default()
	router.GET("/:name/:id", test.BindUri)
	router.Run(":8088")
}

func testCookie() {
	router := gin.Default()
	router.GET("/cookie", test.TestCookie)
	router.Run(":8080")
}

func testLog() {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		fmt.Println("log.....\n")
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage)
	}))
	router.Use(gin.Recovery())
	router.GET("/testing", test.StartPage)
	router.Run(":8080")
}

func testMiddleWare() {
	router := gin.New()
	router.Use(test.Logger())
	router.GET("/testing", test.StartPage)
	router.Run(":8080")
}

func testValidator() {
	router := gin.Default()
	router.Use(test.Logger())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("bookabledate", test.BookableDate)
		if err != nil {
			fmt.Println(err)
		}
	}
	router.GET("/valid", test.GetBookable)
	router.Run(":8080")
}

func testRouterLogFormat() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	router.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})
	router.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	router.Run(":8080")
}

func testGoroutine() {
	router := gin.Default()
	router.GET("/long_async", test.LongAsync)
	router.GET("/long_sync", test.LongSync)
	router.Run(":8080")
}

func testRestart() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(7 * time.Second)
		c.String(http.StatusOK, "hello world")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func testPush() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(test.PushHtml)
	router.GET("/", test.TestPush)
	router.RunTLS(":8080", "./cert/cert.pem", "./cert/key.pem")
}

func testRouterParam() {
	router := gin.Default()
	router.GET("/user/:name", test.TestUserName)
	router.GET("/user/:name/*action", test.TestUserNameAction)
	router.Run(":8080")
}

func testPureJSON() {
	router := gin.Default()

	// 提供 unicode 实体
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}

func testRedirect() {
	router := gin.Default()
	router.GET("/test", test.TestRedirect1)
	router.POST("/test2", test.TestRedirect2)
	router.GET("/test3", func(c *gin.Context) {
		test.TestRedirect3(c, router)
	})
	router.GET("/test4", test.TestRedirect4)
	router.Run(":8080")
}

func testRender() {
	router := gin.Default()
	router.GET("/some_json", test.RenderJSON)
	router.GET("/some_more_json", test.RenderMoreJSON)
	router.GET("/some_xml", test.RenderXML)
	router.GET("/some_yaml", test.RenderYAML)
	router.GET("/some_proto_buf", test.RenderProtoBuf)
	router.Run(":8080")
}

var g errgroup.Group

func testMultiServer() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      test.Router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      test.Router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func testSecureJSON() {
	router := gin.Default()
	router.GET("/names", test.TestSecureJSON)
	router.Run(":8080")
}

func testHeader() {
	router := gin.Default()
	router.Use(test.UseHeader)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}

func testReader() {
	router := gin.Default()
	router.GET("/from_reader", test.TestReader)
	router.Run(":8080")

}

func testLetsEncrypt() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	err := router.RunTLS(":8443", "/Users/mikezhao/certs/server.crt", "/Users/mikezhao/certs/server.key")
	if err != nil {
		panic(err)
	}
}

func testUpload() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8的二进制左移20位 = 8*2^20 ≈ 8M
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./files/" + file.Filename
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "./files/"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}

func testSecret() {
	router := gin.Default()
	test.GetSecretUser(router)
	router.Run(":8080")
}

func testLogFile() {
	gin.DisableConsoleColor()

	if err := os.MkdirAll("./logs", 0755); err != nil {
		panic(err)
	}

	f, _ := os.Create("./logs/test.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8080")
}

func testRouter() {
	router := routers.SetupRouter()
	err := router.Run(":8888")
	if err != nil {
		return
	}
}
