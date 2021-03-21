/*
 * @Author: your name
 * @Date: 2021-03-02 16:39:27
 * @LastEditTime: 2021-03-08 15:48:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gin-blog\service\tag_service\tag.go
 */
package tag_service

import (
	"encoding/json"
	"gin-blog/models"
	"gin-blog/pkg/gredis"
	"gin-blog/pkg/logging"
	"gin-blog/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

// 通过 Name 判断标签名是否已存在
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

// 通过 ID 判断标签是否已存在
func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

// // 添加标签
func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

// 编辑标签
func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

// 获取标签数量
func (t *Tag) Count() (int64, error) {
	return models.GetTagTotal(t.getMaps())
}

// 获取所有标签数据
func (t *Tag) GetAll() ([]models.Tag, error) {
	var tags, cacheTags []models.Tag

	// 先从 Redis 中读取缓存数据
	cache := cache_service.Tag{
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			if err := json.Unmarshal(data, &cacheTags); err != nil {
				logging.Error(err)
			}
			return cacheTags, nil
		}
	}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	// gredis.Set(key, tags, 3600)
	return tags, nil
}

// ...
func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = 0
	if t.Name != "" {
		maps["name"] = t.Name
	}

	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
