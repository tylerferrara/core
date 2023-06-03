package main

type Disability string

const (
	Yes     Disability = "Yes, I Have a Disability, Or Have A History/Record Of Having A Disability"
	No      Disability = "No, I Don't Have A Disability, Or A History/Record Of Having A Disability"
	Decline Disability = "I Don't Wish To Answer"
)

var Disabilities = []Disability{
	Yes,
	No,
	Decline,
}

func (d Disability) isValid() bool {
	for _, disability := range Disabilities {
		if d == disability {
			return true
		}
	}
	return false
}
