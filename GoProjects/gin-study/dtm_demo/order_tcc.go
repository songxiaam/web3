package dtm_demo

import (
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/go-resty/resty/v2"
	"github.com/lithammer/shortuuid/v4"
)

func SubmitOrderTCC() {
	fmt.Println("SubmitOrderTCC")

	err := dtmcli.TccGlobalTransaction(DtmServer, shortuuid.New(), func(tcc *dtmcli.Tcc) (*resty.Response, error) {

		res, reer := tcc.CallBranch(&OrderRequest{ID: 1}, QsOrderServer+"/order", QsOrderServer+"/order_commit", QsOrderServer+"/order_cancel")
		if reer != nil {
			return res, reer
		}

		return tcc.CallBranch(&OrderRequest{ID: 1}, QsOrderServer, QsOrderServer+"/order", QsOrderServer+"/order_cancel")
	})
}
