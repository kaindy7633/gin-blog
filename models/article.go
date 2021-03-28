package models

import "gorm.io/gorm"

type Article struct {
	Model

	TagID int `gorm:"index" json:"tag_id"`
	Tag   Tag `json:"tag"`

	Title         string `gorm:"title" json:"title"`
	Desc          string `gorm:"desc" json:"desc"`
	Content       string `gorm:"content" json:"content"`
	CoverImageUrl string `gorm:"cover_image_url" json:"cover_image_url"`
	CreatedBy     string `gorm:"created_by" json:"created_by"`
	ModifiedBy    string `gorm:"modified_by" json:"modified_by"`
	State         int    `gorm:"state" json:"state"`
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}
