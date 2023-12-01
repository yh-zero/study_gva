package system

import (
	"study_gva/global"
	"study_gva/model/system"
	"study_gva/model/system/request"
)

type DictionaryDetailService struct{}

// // 删除字典详情数据
func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(detail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(&detail).Error
	return err
}

// 更新字典详情数据
func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(detail *system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(detail).Error
	return err
}

// 根据id获取字典详情单条数据
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error

	return
}

// 创建字典详情数据
func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

// 分页获取字典详情列表
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 如果有条件搜索 下发会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}
