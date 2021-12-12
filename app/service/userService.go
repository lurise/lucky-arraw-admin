package service

import "luckyarrow/app/dao"

var UserService =new(userService)

type userService struct {

}

func (*userService) DaligateScore() {
	dao.User.DB().Model("user").Where()
}