package utils

var (
	LoginVerify        = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify     = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	AuthorityVerify    = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify  = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify = Rules{"OldAuthorityId": {NotEmpty()}}
	IdVerify           = Rules{"ID": []string{NotEmpty()}}
	RegisterVerify     = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	ApiVerify          = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify         = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify     = Rules{"Title": {NotEmpty()}}
	AutoPackageVerify  = Rules{"PackageName": {NotEmpty()}}
)
