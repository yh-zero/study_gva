package system

type RouterGroup struct {
	BaseRouter
	InitRouter
	MenuRouter
	UserRouter
	AuthorityRouter
	OperationRecordRouter
	DictionaryRouter
	DictionaryDetailRouter
}
