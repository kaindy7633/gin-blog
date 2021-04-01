package models

import "gorm.io/gorm"

type Article struct {
	Model

	TagID int `gorm:"index" json:"tag_id"`
	Tag   Tag `gorm:"foreignkey:TagID" json:"tag"`

	Title      string `gorm:"title" json:"title"`
	Desc       string `gorm:"desc" json:"desc"`
	Content    string `gorm:"content" json:"content"`
	CreatedBy  string `gorm:"created_by" json:"created_by"`
	ModifiedBy string `gorm:"modified_by" json:"modified_by"`
	State      int    `gorm:"state" json:"state"`
}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_at = ?", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
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

// GetArticle Get a single article based on ID
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_at = ?", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// err = db.Model(&article).Related(&article.Tag).Error
	err = db.Preload("Tag").Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

// AddArticle add a single article
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}

	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

// EditArticle modify a single article
func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND deleted_at = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
