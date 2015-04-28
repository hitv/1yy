package main

import (
	"log"
)

func main() {
	silverAdvert := NewSilverAdvert("http://service.inkey.com", "864093029242286", "15843983987", "123.com..")
	err := silverAdvert.Login()
	if err != nil {
		panic(err)
	}

	//	adverts, err := silverAdvert.IndexAdverts()
	//	if err != nil {
	//		panic(err)
	//	}
	categoryIds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	for _, id := range categoryIds {
		adverts, err := silverAdvert.PullCategoryAds(id)
		if err != nil {
			panic(err)
		}
		log.Printf("拉取分类%d广告，%d条\n", id, len(adverts))

		for _, advert := range adverts {
			if !advert.IsPublicServiceAdvert {
				earn, err := silverAdvert.GeneratedIntegral(advert.Id)
				if err != nil {
					log.Printf("广告(%d)赚钱出错：%s\n", advert.Id, err)
					continue
				}
				log.Printf("广告(%d)赚钱：%d\n", advert.Id, earn)
			}
		}
	}
}
