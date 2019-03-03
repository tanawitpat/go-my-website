package workexperience

type WorkExperienceDB struct {
	CompanyName      string   `bson:"company_name"`
	CompanyImagePath string   `bson:"company_image_path"`
	CompanyLocation  string   `bson:"company_location"`
	StartMonth       string   `bson:"start_month"`
	StartYear        string   `bson:"start_year"`
	EndMonth         string   `bson:"end_month"`
	EndYear          string   `bson:"end_year"`
	IsCurrentWork    bool     `bson:"is_current_work"`
	Position         string   `bson:"position"`
	Description      []string `bson:"description"`
}
