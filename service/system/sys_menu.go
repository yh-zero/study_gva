package system

import (
	"gorm.io/gorm"
	"strconv"

	"study_gva/global"
	"study_gva/model/system"

	"github.com/pkg/errors"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

// 用户角色默认路由检查
func (menuService *MenuService) UserAuthorityDefaultRouter(user *system.SysUser) {
	var menuIds []string
	err := global.GVA_DB.Model(&system.SysAuthorityMenu{}).Where("sys_authority_authority_id = ?", user.AuthorityId).Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am system.SysBaseMenu
	err = global.GVA_DB.First(&am, "name = ? and id in(?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}

// 获取路由总树map
func (menuService *MenuService) getMenuTreeMap(authorityId uint) (treeMap map[string][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	var btns []system.SysAuthorityBtn
	treeMap = make(map[string][]system.SysMenu)

	var SysAuthorityMenys []system.SysAuthorityMenu
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&SysAuthorityMenys).Error
	if err != nil {
		return
	}

	var MenuIds []string
	for i := range SysAuthorityMenys {
		MenuIds = append(MenuIds, SysAuthorityMenys[i].MenuId)
	}

	err = global.GVA_DB.Where("id in (?)", MenuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}

	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}

	err = global.GVA_DB.Where("authority_id = ?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint)
	for _, btn := range btns {
		if btnMap[btn.SysMenuID] == nil {
			btnMap[btn.SysMenuID] = make(map[string]uint)
		}
		btnMap[btn.SysMenuID][btn.SysBaseMenuBtn.Name] = authorityId
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.ID]
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err

}

// 获取动态菜单树
func (menuService *MenuService) GetMenuTree(authorityId uint) (menus []system.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err

}

// 获取子菜单
func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// 为角色添加menu树
func (MenuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, authorityId uint) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}
