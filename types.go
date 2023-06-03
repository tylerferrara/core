package main

import "net/url"

type Pdf struct {
	Content string
}

type Application struct {
	Pdf          Pdf
	PersonalInfo PersonalInfo
	Work         Work
	Identity     Identity
}

type Work struct {
	Experience []Experience
	Education  []Education
	Skills     []Skill
	Websites   []Website
	Socials    []Social
}

type Identity struct {
	requireSponsorship bool
	salary             string
	gender             string
	isHispanicOrLatino bool
	ethnicity          string
	veteranStatus      string
	disabilityStatus   string
}

type Social struct {
	url url.URL
}

type Website struct {
	url url.URL
}

type Skill struct {
	name string
}

type Education struct {
	schoolName   string
	degree       string
	fieldOfStudy string
	gpa          string
	startDate    string
	endDate      string
}

type Experience struct {
	jobTitle         string
	companyName      string
	location         string
	currentlyWorking bool
	startDate        string
	endDate          string
	description      string
}

type PersonalInfo struct {
	firstName string
	lastName  string
	email     string
	phone     string
	zipCode   int32
	address   string
	city      string
	state     string
	county    string
	country   string
}
