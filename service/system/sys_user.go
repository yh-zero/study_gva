package system

import (
	"errors"

	"study_gva/global"
	"study_gva/model/system"
)

type UserService struct{}

// 通过uuid获取用户信息
func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}
