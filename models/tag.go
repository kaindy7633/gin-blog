/*
 * @Description: 标签模块的models
 * @Author: LiuZhen
 * @Date: 2021-02-24 17:13:35
 * @Last Modified by: LiuZhen
 * @Last Modified time: 2021-02-26 17:08:34
 */
package models

type Tag struct {
	Model

	Name       string `gorm:"name" json:"name"`
	CreatedBy  string `gorm:"created_by" json:"created_by"`
	ModifiedBy string `gorm:"modified_by" json:"modified_by"`
	State      int    `gorm:"state" json:"state"`
}

/**
 * @description: 通过名称 name 检查是否重名
 * @param {string} name
 * @return {(bool error)}
 */
// func ExistTagByName(name string) (bool, error) {
// 	var tag Tag
// 	err := db.Select("id").Where("name = ? and deleted_on = ?", name, 0).First(&tag).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return false, err
// 	}

// 	if tag.ID > 0 {
// 		return true, nil
// 	}

// 	return false, nil
// }

// func ExistTagByID(id int) (bool, error) {
// 	var tag Tag
// 	err := db.Select("id").Where("id = ? and deleted_on = ?", id, 0).First(&tag).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return false, err
// 	}

// 	if tag.ID > 0 {
// 		return true, nil
// 	}

// 	return false, nil
// }

/**
 * @description: 添加标签
 * @param {string} name
 * @param {int} state
 * @param {string} createBy
 * @return {error}
 */
// func AddTag(name string, state int, createdBy string) error {
// 	tag := Tag{
// 		Name:      name,
// 		State:     state,
// 		CreatedBy: createdBy,
// 	}

// 	if err := db.Create(&tag).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// 批量获取标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag, err error) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// 获取标签总数量
// func GetTagTotal(maps interface{}) (count int64, err error) {
// 	err = db.Model(&Tag{}).Where(maps).Count(&count).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	return
// }

// // EditTag modify a single tag
// func EditTag(id int, data interface{}) error {
// 	if err := db.Model(&Tag{}).Where("id = ? and deleted_on = ?", id, 0).Updates(data).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
