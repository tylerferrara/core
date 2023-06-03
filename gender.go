package main

type Gender string

const (
	Female      Gender = "Female"
	Male        Gender = "Male"
	NotDeclared Gender = "Not Declared"
)

var Genders = []Gender{
	Female,
	Male,
	NotDeclared,
}

func (g Gender) isValid() bool {
	for _, gender := range Genders {
		if g == gender {
			return true
		}
	}
	return false
}
