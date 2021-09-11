package tools

type Image interface {
	GetUrl(Name string) string
	GetToken() string
}
