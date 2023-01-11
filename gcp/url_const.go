package gcp

const (
	httpsScheme             = "https"
	googleAPIsHost          = "googleapis.com"
	translationAPIHost      = "translation." + googleAPIsHost
	languageTranslateV2Path = "/language/translate/v2"
	languagesPath           = languageTranslateV2Path + "/languages"
	detectPath              = languageTranslateV2Path + "/detect"
)
