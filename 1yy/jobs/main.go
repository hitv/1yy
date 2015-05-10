package main

import (
	"encoding/json"
	"log"

	"hi.tv/1yy/libs/mivideo"
)

//	"hi.tv/1yy/models"

func main() {
	jobsConf, err := NewJobsConfig("./jobs.conf")
	if err != nil {
		panic(err)
	}
	/*
		err = models.InitDB(jobsConf.MySQLDsn)
		if err != nil {
			panic(err)
		}
	*/
	miVideoService := mivideo.NewMiVideoService(jobsConf.MiHost, jobsConf.MiApi, jobsConf.MiKey, jobsConf.MiToken)
	homeData, err := miVideoService.FetchHomeData()
	if err != nil {
		panic(err)
	}
	//	log.Printf("------------ %s -------------\n-------- Slides --------\n", homeData.Hot.Title)
	//	for _, slide := range homeData.Hot.Slides {
	//		log.Printf("%#v\n-----------\n", *slide)
	//	}
	//	log.Printf("\n------------ Channels ------------\n")
	//	for _, channel := range homeData.Hot.Channels {
	//		log.Printf("\n------------ Channel ------------\n%#v\n", *channel)
	//		log.Println("\n------------ Groups ------------\n")
	//		for _, group := range channel.Groups {
	//			log.Printf("\n------------ Group ------------\n%#v\n", *group)
	//			log.Println("\n-------- Items ---------\n")
	//			for _, item := range group.Items {
	//				log.Println("\n-------- Item ---------\n%#v\n", *item)
	//			}
	//			log.Println()
	//		}
	//	}

	jsonBytes, err := json.Marshal(homeData)
	log.Println(string(jsonBytes))
}

/*
func main() {
	jobsConf, err := NewJobsConfig("./jobs.conf")
	if err != nil {
		panic(err)
	}

	err = models.InitDB(jobsConf.MySQLDsn)
	if err != nil {
		panic(err)
	}

	miVideoService := mivideo.NewMiVideoService(jobsConf.MiHost, jobsConf.MiKey, jobsConf.MiToken)
	channels, err := models.ChannelDao.GetAllChannels()
	if err != nil {
		fmt.Printf("ChannelDao.GetAllChannels() error: %s", err)
		return
	}
	for _, channel := range channels {

		var (
			page     = 1
			pageSize = 100
		)

		for {
			fmt.Printf("Start fetch channelId: %d, page: %d\n", channel.MiId, page)
			data, err := miVideoService.FetchMediaFilter(channel.MiId, pageSize, page)
			if err != nil {
				fmt.Printf("miVideoService.FetchMediaFilter(%d, %d, %d) error: %s", channel.MiId, pageSize, page, err)
				continue
			}
			for _, info := range data.Data {
				v := &models.VideoModel{
					Title:             info.MediaName,
					Actors:            info.Actors,
					Director:          info.Director,
					Categorys:         info.Categorys,
					Category:          info.Category,
					Area:              info.Area,
					Description:       info.Desc,
					ShortDescription:  info.ShortDesc,
					Flag:              info.Flag,
					IsMultSet:         info.IsMultSet,
					IssueDate:         info.IssueDate,
					Language:          info.Language,
					LastIssueDate:     info.LastIssueDate,
					Md5:               info.Md5,
					DownloadSrc:       utils.JoinIntsToString(info.DownloadSrc, ","),
					MidType:           info.MidType,
					IndexCapital:      info.NameIndex.Capital,
					IndexFuzzyCapital: info.NameIndex.FuzzyCapital,
					IndexSpelling:     info.NameIndex.Spelling,
					IndexSymbol:       info.NameIndex.Symbol,
					IndexValidLength:  info.NameIndex.ValidLength,
					PlayCount:         info.PlayCount,
					PlayLength:        info.PlayLength,
					PosterURL:         info.PosterURL,
					Resolution:        info.Resolution,
					Score:             info.Score,
					ScoreCount:        info.ScoreCount,
					SetCount:          info.SetCount,
					SetNow:            info.SetNow,
					SmallPosterMd5:    info.SmallPosterMd5,
					SmallPosterURL:    info.SmallPosterURL,
					WebpPosterMd5:     info.WebpPosterMd5,
					WebpPosterURL:     info.WebpPosterURL,
					Status:            enums.StatusEnabled.String(),
					MiId:              info.MediaId,
				}
				var id int64
				video, err := models.VideoDao.GetByMid(info.MediaId)
				if err != nil {
					if err != models.ErrNotExist {
						log.Printf("models.VideoDao.GetByMid(%d) error: %s\n", info.MediaId)
						continue
					}
					id, err = models.VideoDao.Insert(v)
					if err != nil {
						log.Printf("models.VideoDao.Insert(%#v) error: %s\n", video)
						continue
					}
					log.Printf("video.Insert() vid: %d\n", id)
				}
				id, err = models.VideoDao.Update(v)
				if err != nil {
					log.Printf("video.Update() error: %s\n", err)
					continue
				}
				log.Printf("video.Update() vid: %d\n", id)
			}
			if data.Count <= page*pageSize {
				break
			}
			page++
		}
	}
}

/*
	info, err := fetch.FetchChannelInfo(-1)
	if err != nil {
		fmt.Printf("FetchChannelInfo(-1) error: %s\n", err)
		return
	}
	err = info.Save()
	if err != nil {
		fmt.Printf("info.Save() error: %s\n", err)
		return
	}
*/
/*data, err := fetch.FetchMediaDetail(1075038)
if err != nil {
	fmt.Printf("fetch.FetchMediaDetail(%d) error: %s", 1075038, err)
	return
}
fmt.Printf("data: %#v\n\n", data)
*/
