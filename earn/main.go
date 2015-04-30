package main

import (
	"log"
	"time"
)

func main() {
	silverAdvert := NewSilverAdvert("http://service.inkey.com", "864093029225386", "13943969701", "123.com..")
	err := silverAdvert.Login()
	if err != nil {
		panic(err)
	}

	//	adverts, err := silverAdvert.IndexAdverts()
	//	if err != nil {
	//		panic(err)
	//	}
	for id := 32952; id < 53215; id++ {
		go func(id int) {
			earn, err := silverAdvert.GeneratedIntegral(id)
			if err != nil {
				log.Printf("广告(%d)赚钱出错：%s\n", id, err)
				return
			}
			log.Printf("广告(%d)赚钱：%d\n", id, earn)
		}(id)
	}
	<-time.After(time.Minute * 10)
	return
	categoryIds := []int{2, 3, 8}
	for i := 0; i < 30; i++ {
		for _, id := range categoryIds {
			adverts, err := silverAdvert.PullCategoryAds(id)
			if err != nil {
				panic(err)
			}
			log.Printf("拉取分类%d广告，%d条\n", id, len(adverts))

			for _, advert := range adverts {
				if !advert.IsPublicServiceAdvert {
					//id := advert.Id

				}
			}
		}
	}
}
