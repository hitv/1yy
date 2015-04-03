package models

import ()

type MtCategory struct {
	Id          int    //分类id
	Title       string //分类标题
	Description string //分类描述
	Status      string //分类状态
}

func GetCategories() (categories []*MtCategory, err error) {
	key := "AvaliableCategories"
	data, exists := cache.Get(key)
	if exists {
		categories = data.([]*MtCategory)
		return
	}
	rows, err := db.Query("SELECT id,title,description,status FROM category WHERE status='NORMAL' ORDER BY `order` ASC")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		category := &MtCategory{}
		err = rows.Scan(&category.Id, &category.Title, &category.Description, &category.Status)
		if err != nil {
			return
		}
		categories = append(categories, category)
	}
	cache.Set(key, categories, 0)
	return
}
