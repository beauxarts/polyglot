package polyglot

type Translator interface {
	Languages(language string) (map[string]string, error)
	Detect(content string) (string, error)
	Translate(source, target string, format TranslateFormat, query ...string) ([]string, error)
	IsHTMLSupported() bool
}
