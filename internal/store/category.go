package store

type Category struct {
	ID               string     `gorm:"primaryKey" json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	ParentCategoryID *string    `json:"parent_category_id"`
	Children         []Category `gorm:"foreignKey:ParentCategoryID" json:"children,omitempty"`
}
