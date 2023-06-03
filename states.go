package main

type State string

const (
	Alabama                          State = "Alabama"
	Alaska                           State = "Alaska"
	AmericanSamoa                    State = "American Samoa"
	Arizona                          State = "Arizona"
	Arkansas                         State = "Arkansas"
	ArmedForcesAmericas              State = "Armed Forces Americas"
	ArmedForcesEurope                State = "Armed Forces Europe"
	ArmedForcesPacific               State = "Armed Forces Pacific"
	California                       State = "California"
	Colorado                         State = "Colorado"
	Connecticut                      State = "Connecticut"
	Delaware                         State = "Delaware"
	DistrictOfColumbia               State = "District Of Columbia"
	Florida                          State = "Florida"
	Georgia                          State = "Georgia"
	Guam                             State = "Guam"
	Hawaii                           State = "Hawaii"
	Idaho                            State = "Idaho"
	Illinois                         State = "Illinois"
	Indiana                          State = "Indiana"
	Iowa                             State = "Iowa"
	Kansas                           State = "Kansas"
	Kentucky                         State = "Kentucky"
	Louisiana                        State = "Louisiana"
	Maine                            State = "Maine"
	Maryland                         State = "Maryland"
	Massachusetts                    State = "Massachusetts"
	Michigan                         State = "Michigan"
	Minnesota                        State = "Minnesota"
	Mississippi                      State = "Mississippi"
	Missouri                         State = "Missouri"
	Montana                          State = "Montana"
	Nebraska                         State = "Nebraska"
	Nevada                           State = "Nevada"
	NewHampshire                     State = "New Hampshire"
	NewJersey                        State = "New Jersey"
	NewMexico                        State = "New Mexico"
	NewYork                          State = "New York"
	NorthCarolina                    State = "North Carolina"
	NorthDakota                      State = "North Dakota"
	NorthernMarianaIslands           State = "Northern Mariana Islands"
	Ohio                             State = "Ohio"
	Oklahoma                         State = "Oklahoma"
	Oregon                           State = "Oregon"
	Pennsylvania                     State = "Pennsylvania"
	PuertoRico                       State = "Puerto Rico"
	RhodeIsland                      State = "Rhode Island"
	SouthCarolina                    State = "South Carolina"
	SouthDakota                      State = "South Dakota"
	Tennessee                        State = "Tennessee"
	Texas                            State = "Texas"
	UnitedStatesMinorOutlyingIslands State = "United States Minor Outlying Islands"
	Utah                             State = "Utah"
	Vermont                          State = "Vermont"
	VirginIslands                    State = "Virgin Islands"
	Virginia                         State = "Virginia"
	Washington                       State = "Washington"
	WestVirginia                     State = "West Virginia"
	Wisconsin                        State = "Wisconsin"
	Wyoming                          State = "Wyoming"
)

var States = []State{
	Alabama,
	Alaska,
	AmericanSamoa,
	Arizona,
	Arkansas,
	ArmedForcesAmericas,
	ArmedForcesEurope,
	ArmedForcesPacific,
	California,
	Colorado,
	Connecticut,
	Delaware,
	DistrictOfColumbia,
	Florida,
	Georgia,
	Guam,
	Hawaii,
	Idaho,
	Illinois,
	Indiana,
	Iowa,
	Kansas,
	Kentucky,
	Louisiana,
	Maine,
	Maryland,
	Massachusetts,
	Michigan,
	Minnesota,
	Mississippi,
	Missouri,
	Montana,
	Nebraska,
	Nevada,
	NewHampshire,
	NewJersey,
	NewMexico,
	NewYork,
	NorthCarolina,
	NorthDakota,
	NorthernMarianaIslands,
	Ohio,
	Oklahoma,
	Oregon,
	Pennsylvania,
	PuertoRico,
	RhodeIsland,
	SouthCarolina,
	SouthDakota,
	Tennessee,
	Texas,
	UnitedStatesMinorOutlyingIslands,
	Utah,
	Vermont,
	VirginIslands,
	Virginia,
	Washington,
	WestVirginia,
	Wisconsin,
	Wyoming,
}

func (s State) IsValid() bool {
	for _, state := range States {
		if s == state {
			return true
		}
	}
	return false
}
