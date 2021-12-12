package service

import (
	"github.com/gogf/gf/frame/g"
	"luckyarrow/app/model"
)

var PrizeService = new(prizeService)

type prizeService struct{}

func (*prizeService) GetPrizeList() (prizes []*model.Prize, err error) {
	err = g.DB().Model("prize").Where("status = ?", 2).Scan(&prizes)
	return prizes, err
}
