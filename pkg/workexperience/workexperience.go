package workexperience

import (
	"fmt"
	"go-my-resume/app"
	"net/http"
)

func InquiryWorkExperience(w http.ResponseWriter, req *http.Request) {
	session, err := app.GetMongoSession()
	if err != nil {
		fmt.Println("Cannot get mongo session", err)
	}
	workExperience := []WorkExperienceDB{}
	defer session.Close()
	db := session.DB("myresume")
	db.C("work_experience").Find(nil).All(&workExperience)
	fmt.Println(workExperience)
	tpl.ExecuteTemplate(w, "workexperience.gohtml", nil)
}
