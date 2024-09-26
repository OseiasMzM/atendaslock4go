package model

type License struct {
	name  string
	count int
}

func NewLicense(name string, count int) License {
	return License{name, count}
}
