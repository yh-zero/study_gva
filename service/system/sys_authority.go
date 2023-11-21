package system

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"

	"study_gva/global"
	"study_gva/model/common/request"
	"study_gva/model/system"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

// 分页获取数据
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize // 每页大小
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAuthority{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id", "0").Find(&authority).Error
	for k := range authority {
		err = authorityService.findChildrenAuthority(&authority[k])
	}
	return authority, total, err

}

// 查询子角色
func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

// 创建一个角色
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	err = global.GVA_DB.Create(&auth).Error
	return auth, err

}

// 菜单与角色绑定
func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First("authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Where(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// 删除角色
func (authority *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if errors.Is(global.GVA_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	fmt.Println("DeleteAuthority: auth.AuthorityId", auth)
	fmt.Println("DeleteAuthority: auth.AuthorityId", auth.AuthorityId)

	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不运行删除")
	}
	db := global.GVA_DB.Preload("SysBaseMenus").Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		err = global.GVA_DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
	}
	if len(auth.DataAuthorityId) > 0 {
		err = global.GVA_DB.Model(auth).Association("DataAuthorityId").Delete(auth.DataAuthorityId)
		if err != nil {
			return
		}
	}
	err = global.GVA_DB.Delete(&[]system.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ?", auth.AuthorityId).Error
	if err != nil {
		return
	}
	authorityId := strconv.Itoa(int(auth.AuthorityId))
	CasbinServiceApp.ClearCasbin(0, authorityId)
	return err

}

// 更新一个角色对象
func (AuthorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Updates(&auth).Error
	return auth, err
}
