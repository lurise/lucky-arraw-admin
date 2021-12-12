package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"luckyarrow/app/service"
)

var PrizeApi = new(prizeApi)

type prizeApi struct{}

func (*prizeApi) GetPrizeList(r *ghttp.Request) {
	prizes, err := service.PrizeService.GetPrizeList()
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
	}
	r.Response.WriteJson(g.Map{
		"code":200,
		"data":prizes,
		"msg":"success to get the prizes.",
	})
}
