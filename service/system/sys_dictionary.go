package system

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"study_gva/global"
	"study_gva/model/system"
)

type DictionaryService struct{}

// 分页获取字典列表
func (dictionary *DictionaryService) GetSysDictionaryInfoList() (list interface{}, err error) {
	var sysDictionarys []system.SysDictionary
	err = global.GVA_DB.Find(&sysDictionarys).Error
	return sysDictionarys, err
}

func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type,不运行创建")
	}
	err = global.GVA_DB.Create(&sysDictionary).Error
	return err
}
