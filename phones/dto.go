package phones

// Country model
type Country struct {
	Code   string
	Name   string
	RegExp string `json:"-"`
}

// Phone model
type Phone struct {
	ID          int
	Name        string
	Phone       string
	CountryCode string
	CountryName string
	State       string
}