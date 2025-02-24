package models

import "gorm.io/gorm"

type Song struct {
    gorm.Model
    Group       string `json:"group"`
    Song        string `json:"song"`
    ReleaseDate string `json:"releaseDate"`
    Text        string `json:"text"`
    Link        string `json:"link"`
}
