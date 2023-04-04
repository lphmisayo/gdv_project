package models

type CommentModel struct {
	MODEL

	SubComments        []*CommentModel `json:"sub_comments" gorm:"foreignKey:ParentCommentID"`
	ParentCommentModel *CommentModel   `json:"parent_comment_model" gorm:"foreignKey:ParentCommentID"`
	ParentCommentID    *uint           `json:"parent_comment_id"`
	Content            string          `json:"content" gorm:"size:256"`
	DiggCount          int             `json:"digg_count" gorm:"size:8;default:0"`
	CommentCount       int             `json:"comment_count" gorm:"size:8;default:0"`
	Article            ArticleModel    `json:"-" gorm:"foreignKey:ArticleID"`
	ArticleID          uint            `json:"article_id"`
	User               UserModel       `json:"user"`
	UserID             uint            `json:"user_id"`
}
