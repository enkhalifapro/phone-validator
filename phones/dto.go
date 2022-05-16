package phones

type Country struct {
	Code   string `json:"code"`
	Name   string `json:"country"`
	RegExp string `json:"-"`
}
