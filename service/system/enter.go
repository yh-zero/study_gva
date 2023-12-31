package system

type ServiceGroup struct {
	JwtService
	UserService
	CasbinService
	OperationRecordService
	InitDBService
	MenuService
	AuthorityService
	DictionaryService
	DictionaryDetailService
	ApiService
	BaseMenuService
	AutoCodeService
}
