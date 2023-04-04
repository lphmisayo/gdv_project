package models

type MessageModel struct {
	MODEL
	SendUserID       uint      `json:"send_user_id" gorm:"primaryKey"`
	SendUserModel    UserModel `json:"-" gorm:"foreignKey:SendUserID"`
	SendUserNickName string    `json:"send_user_nick_name" gorm:"size:42"`
	SendUserAvatar   string    `json:"send_user_avatar"`

	RevUserID       uint      `json:"rev_user_id" gorm:"primaryKey"`
	RevUserModel    UserModel `json:"-" gorm:"foreignKey:RevUserID"`
	RevUserNickName string    `json:"rev_user_nick_name" gorm:"size42"`
	RevUserAvatar   string    `json:"rev_user_avatar"`

	IsRead  bool   `json:"is_read" gorm:"default:false"`
	Content string `json:"content"`
}
