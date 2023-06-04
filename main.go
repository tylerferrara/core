package main

// This takes a LONG time to complete
func main() {
	q := ask("Im going to give you a text block of the ASCII strings taken from a candidates resume. Your objective is to fill the values for a JSON object structure. We will present you the JSON object, prompting you to collect the correct values for each key in the JSON object. Return the complete JSON object without any extra words or explination. We are looking for correctness. Understand?")
	run := ask("Here is the PDF of the resume: " + collect("test.pdf"))
	win := ask(`The JSON structure is this:
	{
		"pdf": {
		  "content": ""
		},
		"personal_info": {
		  "first_name": "",
		  "last_name": "",
		  "email": "",
		  "phone": "",
		  "zip_code": 0,
		  "address": "",
		  "city": "",
		  "state": "",
		  "county": "",
		  "country": ""
		},
		"work": {
		  "experience": [],
		  "education": [],
		  "skills": [],
		  "websites": [],
		  "socials": []
		},
		"identity": {
		  "require_sponsorship": false,
		  "salary": "",
		  "gender": "",
		  "is_hispanic_or_latino": false,
		  "ethnicity": "",
		  "veteran_status": "",
		  "disability_status": ""
		}
	  }`)
	continueSession := ContinuousSession{
		sessions: []Session{q, run, win},
	}
	launch(continueSession.sessions)
}
