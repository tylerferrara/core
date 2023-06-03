package main

type Ethnicity string

const (
	AmericanIndian         Ethnicity = "American Indian or Alaska Native"
	Asian                  Ethnicity = "Asian"
	BlackOrAfricanAmerican Ethnicity = "Black or African American"
	HispanicOrLatino       Ethnicity = "Hispanic or Latino"
	NoAnswer               Ethnicity = "I do not wish to answer"
	PacificIslander        Ethnicity = "Native Hawaiian or Other Pacific Islander"
	TwoOrMore              Ethnicity = "Two or More Races"
	White                  Ethnicity = "White"
)

var Ethnicities = []Ethnicity{
	AmericanIndian,
	Asian,
	BlackOrAfricanAmerican,
	HispanicOrLatino,
	NoAnswer,
	PacificIslander,
	TwoOrMore,
	White,
}

func (e Ethnicity) isValid() bool {
	for _, ethnicity := range Ethnicities {
		if e == ethnicity {
			return true
		}
	}
	return false
}
