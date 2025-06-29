package dtm_demo

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"gorm.io/driver/mysql"
	"net/http"
)

func orderMsgSuccess(c *gin.Context) {
	fmt.Println("orderSuccess")

	c.JSON(http.StatusOK, gin.H{})
}
func orderMsgFail(c *gin.Context) {
	fmt.Println("orderFail")
	c.JSON(http.StatusOK, gin.H{})

}
func stockMsgSuccess(c *gin.Context) {
	fmt.Println("stockSuccess")
	//c.JSON(http.StatusNotFound, gin.H{})
	c.AbortWithStatus(http.StatusInternalServerError)
	return
}
func stockMsgFail(c *gin.Context) {
	fmt.Println("stockFail")
	c.JSON(http.StatusOK, gin.H{})
	//c.AbortWithStatus(http.StatusInternalServerError)
}

func SubmitOrderMsg() {
	fmt.Println("SubmitOrderMsg")
	req := OrderRequest{
		ID: 2,
	}
	// 注册,表示本地事务成功后,调用接口
	msg := dtmcli.NewMsg(DtmServer, shortuuid.New()).
		Add(QsOrderServer+"/order", req)
	msg.DoAndSubmitDB(QsOrderServer + "/")

}
