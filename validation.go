package main

import (
	"time"
)

func (app Application) IsValid() bool {
	return (app.Pdf.Content != "" &&
		app.PersonalInfo.isValid() &&
		app.Work.IsValid() &&
		app.Identity.IsValid())
}

func (work Work) IsValid() bool {
	for _, edu := range work.Education {
		if !edu.IsValid() {
			return false
		}
	}
	for _, exp := range work.Experience {
		if !exp.IsValid() {
			return false
		}
	}
	return true
}

func (identity Identity) IsValid() bool {
	return (identity.salary != "" &&
		Gender(identity.gender).isValid() &&
		Ethnicity(identity.ethnicity).isValid() &&
		Veteran(identity.veteranStatus).isValid() &&
		Disability(identity.disabilityStatus).isValid())
}

func (edu Education) IsValid() bool {
	return (edu.schoolName != "" &&
		Degree(edu.degree).IsValid() &&
		edu.fieldOfStudy != "" &&
		edu.gpa != "" &&
		edu.startDate != "" &&
		edu.endDate != "" &&
		IsValidDate(edu.startDate) &&
		IsValidDate(edu.endDate))
}

func (exp Experience) IsValid() bool {
	return (exp.jobTitle != "" &&
		exp.companyName != "" &&
		exp.location != "" &&
		exp.description != "" &&
		exp.startDate != "" &&
		exp.endDate != "" &&
		IsValidDate(exp.startDate) &&
		IsValidDate(exp.endDate))
}

func (info PersonalInfo) isValid() bool {
	return (info.firstName != "" &&
		info.lastName != "" &&
		info.email != "" &&
		info.phone != "" &&
		info.address != "" &&
		info.zipCode >= 00501 &&
		info.zipCode <= 99950 &&
		info.city != "" &&
		State(info.state).IsValid() &&
		info.country != "" &&
		info.county != "")
}

func IsValidDate(s string) bool {
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}
