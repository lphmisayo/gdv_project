package models

type FadeBackModel struct {
	MODEL
	Email        string `json:"email" gorm:"size:64"`
	Content      string `json:"content" gorm:"size:128"`
	ApplyContent string `json:"applyContent" gorm:"size:128"`
	IsApply      bool   `json:"isApply"`
}
