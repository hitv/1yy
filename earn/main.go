package main

import (
	"log"
	"strings"
	"time"
)

var accounts = [][]string{
	{"864093321394329", "15843983987", "123.com.."},
	{"864093029225386", "13943969701", "123.com.."},
}

func earn(account []string) {
	silverAdvert := NewSilverAdvert("http://service.inkey.com", account[0], account[1], account[2])
	err := silverAdvert.Login()
	if err != nil {
		log.Printf("%s登录失败\n", account[1])
		return
	}

	//	adverts, err := silverAdvert.IndexAdverts()
	//	if err != nil {
	//		panic(err)
	//	}
	/*for id := 32952; id < 53215; id++ {
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
	return*/
	for n := 0; n < 4; n++ {
		for id := 1; id < 16; id++ {
			log.Printf("开始拉取分类%d广告\n", id)
			adverts, err := silverAdvert.PullCategoryAds(id)
			if err != nil {
				panic(err)
			}
			log.Printf("完成拉取分类%d广告，%d条\n", id, len(adverts))

			for _, advert := range adverts {
				if !advert.IsPublicServiceAdvert {
					earn, err := silverAdvert.GeneratedIntegral(advert.Id)
					if err != nil {
						if strings.Contains(err.Error(), "捡满了") {
							return
						}
						log.Printf("广告(%d)赚钱出错：%s\n", advert.Id, err)
						continue
					}
					log.Printf("广告(%d)赚钱：%d\n", advert.Id, earn)
				}
				time.Sleep(time.Second * 5)
			}
		}
	}
}

func main() {
	for _, account := range accounts {
		earn(account)
	}
}
