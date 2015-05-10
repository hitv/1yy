package controllers

type Index struct {
	*Base
}

func (c *Index) Index() {
	c.Render.HTML(200, "index", map[string]interface{}{
		"channels": []map[string]interface{}{
			map[string]interface{}{
				"more":  "更多热门视频",
				"total": "210741个",
				"data": []map[string]interface{}{
					{
						"title": "新闻",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					}, {
						"title": "娱乐",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					}, {
						"title": "体育",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					}, {
						"title": "音乐",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"more":  "进入电视剧频道",
				"total": "4336部",
				"data": []map[string]interface{}{
					{
						"title": "热播",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					}, {
						"title": "卫视同步",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					}, {
						"title": "美剧",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					}, {
						"title": "日韩",
						"videos": []map[string]interface{}{
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池",
								"thumb": "img/1.jpg",
								"info":  "00:54",
							},
							{
								"title": "赵薇自曝曾“虐”女儿小四月 把女儿推入泳池2",
								"thumb": "img/2.jpg",
								"info":  "00:53",
							},
						},
					},
				},
			},
		},
	})
}

func (c *Index) Best() {
	c.Render.HTML(200, "best", nil)
}

func (c *Index) Rank() {
	c.Render.JSON(200, "rank")
}
