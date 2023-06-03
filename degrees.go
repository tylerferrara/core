package main

type Degree string

const (
	GED        Degree = "GED"
	HighSchool Degree = "High School"
	Associates Degree = "Associates"
	Bachelors  Degree = "Bachelors"
	Masters    Degree = "Masters"
	Doctorate  Degree = "Doctorate"
	BA         Degree = "BA - Double Major"
	MA         Degree = "MA - Double Major"
)

var Degrees = []Degree{
	GED,
	HighSchool,
	Associates,
	Bachelors,
	Masters,
	Doctorate,
	BA,
	MA,
}

func (d Degree) IsValid() bool {
	for _, degree := range Degrees {
		if d == degree {
			return true
		}
	}
	return false
}
