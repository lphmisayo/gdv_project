package models

type TagModel struct {
	MODEL
	Title        string         `gorm:"size:16"`
	ArticleModel []ArticleModel `gorm:"many2many:article_tag" json:"-"`
}
