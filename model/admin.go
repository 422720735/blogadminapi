package model

type SystemConfig struct {
	Id int `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SystemConfigFilters struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Keywords string `json:"keywords"`
	Description string `json:"description"`
	Email string `json:"email"`
	Start string `json:"start"`
	Qq string `json:"qq"`
}

type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
}