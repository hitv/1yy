package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type Account struct {
	Imei string
	Username string
	Password string
}
func getAdverts(silverAdvert *SilverAdvert, categaryIds []int) (adverts []Advert, err error) {
	for id := 1; id < 16; id++ {
		time.Sleep(time.Second * 5)
		tmpAdverts, err := silverAdvert.PullCategoryAds(id)
		if err != nil {
			log.Printf("category id: %d, error: %s\n", id, err)
			continue
		}
		log.Printf("category id: %d, num: %d", id, len(tmpAdverts))
		adverts = append(adverts, tmpAdverts...)
	}
	return
}

func earn(account *Account) {
	silverAdvert := NewSilverAdvert("http://service.inkey.com", account.Imei, account.Username, account.Password)
	err := silverAdvert.Login()
	if err != nil {
		log.Printf("%s login error:%s\n", account.Username, err)
		return
	}
	time.Sleep(time.Second * 5)

	categaryIds := []int{3, 8}
	adverts, err := getAdverts(silverAdvert, categaryIds)
	if err!= nil {
		log.Printf("%s getAdverts error:%s\n", account.Username, err)
		return
	}
	excludesAdverts := make(map[int]bool)
	for n := 0; n < 30; n++ {
		for _, advert := range adverts {
			if _, ok := excludesAdverts[advert.Id]; !ok && !advert.IsPublicServiceAdvert {
				time.Sleep(time.Second * 5)
				earn, err := silverAdvert.GeneratedIntegral(advert.Id)
				if err != nil {
					if strings.Contains(err.Error(), "捡满了") {
						return
					}
					if strings.Contains(err.Error(), "最大播放数") {
						excludesAdverts[advert.Id] = true
						continue
					}

					log.Printf("ad(%d) error: %s\n", advert.Id, err)
					continue
				}
				log.Printf("ad(%d) earn: %d\n", advert.Id, earn)
			}
		}
		time.Sleep(time.Second * 5)

	}
}

func main() {
	if len(os.Args) < 4 {
		log.Printf("Usage: %s imei username password\n", os.Args[0])
		return
	}
	account := &Account{os.Args[1], os.Args[2], os.Args[3]}
	earn(account)
}
