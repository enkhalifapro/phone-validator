package phones

type Country struct {
	Code   string `json:"code"`
	Name   string `json:"country"`
	RegExp string `json:"-"`
}

type Phone struct {
	ID int
	Name string
	Phone string
	CountryCode string
	CountryName string
}