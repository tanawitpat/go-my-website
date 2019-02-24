package workexperience

type WorkExperienceDB struct {
	Company       string   `bson:"company"`
	StartMonth    string   `bson:"start_month"`
	StartYear     string   `bson:"start_year"`
	EndMonth      string   `bson:"end_month"`
	EndYear       string   `bson:"end_year"`
	IsCurrentWork string   `bson:"is_current_work"`
	Position      string   `bson:"position"`
	Description   []string `bson:"description"`
}
