package entity

type Recipe struct {
	Id        int             `gorm:"primaryKey" json:"id"`
	Name      string          `gorm:"not null" json:"name"`
	Dose      string          `gorm:"not null" json:"dose"`
	Type      string          `gorm:"not null" json:"type"`
	PatientId int             `json:"patient_id"`
	Patient   PatientResponse `json:"patient"`
	Tags      []Tag           `gorm:"many2many:recipe_tags" json:"tags"`
	TagsId    []int           `gorm:"-" json:"tags_id"`
}

type RecipeResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Dose      string `json:"dose"`
	Type      string `json:"type"`
	PatientId int    `json:"-"`
}

type RecipeResponseTag struct {
	Id        int             `json:"id"`
	Name      string          `json:"name"`
	Dose      string          `json:"dose"`
	Type      string          `json:"type"`
	PatientId int             `json:"-"`
	Patient   PatientResponse `json:"patient"`
	Tags      []Tag           `json:"tags" gorm:"many2many:recipe_tags;ForeignKey:Id;joinForeignKey:RecipeId;References:Id;joinReferences:TagId"`
}

func (RecipeResponse) TableName() string {
	return "recipes"
}

func (RecipeResponseTag) TableName() string {
	return "recipes"
}
