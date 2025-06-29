package dtm_demo

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"log"
	"net/http"
)

type OrderRequest struct {
	ID int `json:"id"`
}

func InitOrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/api/qsorder")
	{
		orderGroup.POST("/order", orderSuccess)
		orderGroup.POST("/order_compensate", orderFail)
		orderGroup.POST("/stock", stockSuccess)
		orderGroup.POST("/stock_compensate", stockFail)

	}

}

func orderSuccess(c *gin.Context) {
	fmt.Println("orderSuccess")
	c.JSON(http.StatusOK, gin.H{})
}
func orderFail(c *gin.Context) {
	fmt.Println("orderFail")
	c.JSON(http.StatusOK, gin.H{})

}
func stockSuccess(c *gin.Context) {
	fmt.Println("stockSuccess")
	//c.JSON(http.StatusNotFound, gin.H{})
	c.AbortWithStatus(http.StatusInternalServerError)
	return
}
func stockFail(c *gin.Context) {
	fmt.Println("stockFail")
	c.JSON(http.StatusOK, gin.H{})
	//c.AbortWithStatus(http.StatusInternalServerError)
}

func SubmitOrder() {
	fmt.Println("SubmitOrder")
	//req := &gin.H{"id": 1}
	//req := map[string]interface{}{"id": 1}
	req := OrderRequest{
		ID: 2,
	}
	order := QsOrderServer + "/order"
	orderFail := QsOrderServer + "/order_compensate"
	stock := QsOrderServer + "/stock"
	stockFail := QsOrderServer + "/stock_compensate"

	fmt.Println(order)
	fmt.Println(orderFail)
	fmt.Println(stock)
	fmt.Println(stockFail)

	//saga := dtmcli.NewSaga(DtmServer, shortuuid.New()).
	//	Add(
	//		QsOrderServer+"/order", QsOrderServer+"/order_compensate", req).
	//	Add(QsOrderServer+"/stock", QsOrderServer+"/stock_compensate", req)

	saga := dtmcli.NewSaga(DtmServer, shortuuid.New()).
		Add(
			QsOrderServer+"/order", QsOrderServer+"/order_compensate", req).
		Add(QsOrderServer+"/stock", QsOrderServer+"/stock_compensate", req)
	saga.RetryInterval = 2
	saga.TimeoutToFail = 5
	err := saga.Submit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SubmitOrder 已提交")
}
