package workexperience

import (
	"fmt"

	"github.com/globalsign/mgo"
)

func InquiryWorkExperience(session *mgo.Session) ([]WorkExperienceDB, error) {
	workExperience := []WorkExperienceDB{}
	db := session.DB("myresume")
	if err := db.C("work_experience").Find(nil).All(&workExperience); err != nil {
		fmt.Println("Cannot get work experience from the database: ", err)
		return workExperience, err
	}
	fmt.Println(workExperience)
	return workExperience, nil
}
