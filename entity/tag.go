package entity

type Tag struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type TagResponse struct {
	Id     int              `json:"id"`
	Name   string           `json:"name"`
	Recipe []RecipeResponse `json:"recipes" gorm:"many2many:recipe_tags;ForeignKey:Id;joinForeignKey:TagId;References:Id;joinReferences:RecipeId"`
}

func (TagResponse) TableName() string {
	return "tags"
}
