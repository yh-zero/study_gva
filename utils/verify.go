package utils

var (
	LoginVerify        = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify     = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	AuthorityVerify    = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify  = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify = Rules{"OldAuthorityId": {NotEmpty()}}
	IdVerify           = Rules{"ID": []string{NotEmpty()}}
)
