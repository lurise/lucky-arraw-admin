package api

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"luckyarrow/app/model"
	"net/http"
)

var (
	QQApi = new(qqApi)
)

const QQ_CODE2SESSION_URL = "https://api.q.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

type qqApi struct {
}

func (*qqApi) GetUnionId(r *ghttp.Request) {
	var qqResponse model.QQApiResponse

	code := r.GetString("code")
	url := fmt.Sprintf(QQ_CODE2SESSION_URL, "1112078237", "zed1tNV7vPKY0tXb", code)

	resp, err := http.Get(url)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"err":  err.Error(),
		})
		return
	}

	buffer, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buffer, &qqResponse)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"err":  err.Error(),
		})
		return
	}
	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "success to get the unionid",
		"data": qqResponse,
	})
}
