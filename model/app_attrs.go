package model

type AppAttrs struct {
	AppAttributes
	LangKey  string `json:"lang_key"`
	LangName string `json:"lang_name"`
	LangIcon string `json:"lang_icon"`
}
