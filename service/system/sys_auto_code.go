package system

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"study_gva/global"
	"study_gva/model/system"
	"study_gva/resource/autocode_template/subcontract"
	"study_gva/utils"
	ast2 "study_gva/utils/ast"
	"text/template"
)

const (
	packageService     = "service/%s/enter.go"
	packageServiceName = "service"
	packageRouter      = "router/%s/enter.go"
	packageRouterName  = "router"
	packageAPI         = "api/v1/%s/enter.go"
	packageAPIName     = "api/v1"
)

type autoPackage struct {
	path string
	temp string
	name string
}

var (
	injectionPaths      []injectionMeta
	packageInjectionMap map[string]astInjectionMeta
)

func Init(Package string) {
	global.GVA_LOG.Info("Init ---!")
	injectionPaths = []injectionMeta{
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SApi, Package), "enter.go"),
			funcName:    "ApiGroup",
			structNameF: "%sApi",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SRouter, Package), "enter.go"),
			funcName:    "RouterGroup",
			structNameF: "%sRouter",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SService, Package), "enter.go"),
			funcName:    "ServiceGroup",
			structNameF: "%sService",
		},
	}

	packageInjectionMap = map[string]astInjectionMeta{
		packageServiceName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "service", "enter.go"),
			importCodeF:  "study_gva/%s/%s",
			packageNameF: "%s",
			groupName:    "ServiceGroup",
			structNameF:  "%sServiceGroup",
		},
		packageRouterName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "router", "enter.go"),
			importCodeF:  "study_gva/%s/%s",
			packageNameF: "%s",
			groupName:    "RouterGroup",
			structNameF:  "%s",
		},
		packageAPIName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "api/v1", "enter.go"),
			importCodeF:  "study_gva/%s/%s",
			packageNameF: "%s",
			groupName:    "ApiGroup",
			structNameF:  "%sApiGroup",
		},
	}
}

type (
	injectionMeta struct {
		path        string
		funcName    string
		structNameF string // 带格式化的
	}

	astInjectionMeta struct {
		path         string
		importCodeF  string
		structNameF  string
		packageNameF string
		groupName    string
	}
)
type AutoCodeService struct{}

var AutoCodeServiceApp = new(AutoCodeService)

func (autoCodeService *AutoCodeService) DelPackage(autoCode system.SysAutoCode) error {
	return global.GVA_DB.Delete(&autoCode).Error
}

func (autoCOdeService *AutoCodeService) CreateAutoCode(autoCode *system.SysAutoCode) error {
	if autoCode.PackageName == "autocode" || autoCode.PackageName == "system" || autoCode.PackageName == "example" || autoCode.PackageName == "" {
		return errors.New("不能使用已保留的package name")
	}
	if !errors.Is(global.GVA_DB.Where("package_name = ?", autoCode.PackageName).First(&system.SysAutoCode{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同PackageName")
	}

	if err := autoCOdeService.CreatePackageTemp(autoCode.PackageName); err != nil {
		return err
	}
	return global.GVA_DB.Create(&autoCode).Error
}

func (autoCodeServcie *AutoCodeService) CreatePackageTemp(packageName string) error {
	global.GVA_LOG.Info("CreatePackageTemp ---!")
	Init(packageName)
	pendingTemp := []autoPackage{{
		path: packageService,
		name: packageServiceName,
		temp: string(subcontract.Server),
	}, {
		path: packageRouter,
		name: packageRouterName,
		temp: string(subcontract.Router),
	}, {
		path: packageAPI,
		name: packageAPIName,
		temp: string(subcontract.API),
	}}
	fmt.Println("pendingTemp1 ==:", pendingTemp)
	for i, s := range pendingTemp {
		pendingTemp[i].path = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, filepath.Clean(fmt.Sprintf(s.path, packageName)))
	}
	fmt.Println("pendingTemp2 ==:", pendingTemp)

	// 选择模板
	for _, s := range pendingTemp {
		err := os.MkdirAll(filepath.Dir(s.path), 0755)
		if err != nil {
			return err
		}
		f, err := os.Create(s.path)
		if err != nil {
			return err
		}

		defer f.Close()

		temp, err := template.New("").Parse(s.temp)
		if err != nil {
			return err
		}
		err = temp.Execute(f, struct {
			PackageName string `json:"package_name"`
		}{packageName})
		if err != nil {
			return err
		}
	}
	// 创建完成后在对应的位置插入结构代码
	for _, v := range pendingTemp {
		meta := packageInjectionMap[v.name]
		err := ast2.ImportReference(meta.path, fmt.Sprintf(meta.importCodeF, v.name, packageName), fmt.Sprintf(meta.structNameF, utils.FirstUpper(packageName)), fmt.Sprintf(meta.packageNameF, packageName), meta.groupName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (autoCodeService *AutoCodeService) GetPackage() (pkgList []system.SysAutoCode, err error) {
	err = global.GVA_DB.Find(&pkgList).Error
	return pkgList, err
}
