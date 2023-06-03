package main

type Veteran string

const (
	MoreThanOneVeteran  Veteran = "I Identify as one or more of the classifications of protected veterans listed above."
	NonProtectedVeteran Veteran = "I Identify as a Veteran, just not a protected Veteran"
	NotAVeteran         Veteran = "I am not a Veteran"
	DeclineToAnswer     Veteran = "Decline to Specify."
)

var Veterans = []Veteran{
	MoreThanOneVeteran,
	NonProtectedVeteran,
	NotAVeteran,
	DeclineToAnswer,
}

func (v Veteran) isValid() bool {
	for _, vet := range Veterans {
		if v == vet {
			return true
		}
	}
	return false
}
