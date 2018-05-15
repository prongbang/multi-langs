package model

type AppLangs struct {
	AppLanguage
	Key  string `json:"key"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
