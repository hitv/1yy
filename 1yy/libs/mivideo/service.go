package mivideo

import (
	"io"
	"net/http"

	"github.com/bitly/go-simplejson"
	"hi.tv/1yy/libs/mivideo/errors"
	"hi.tv/1yy/libs/mivideo/types"
)

type MiVideoService interface {
	NewRequest(path string) *Request
	FetchHomeData() (*types.HomeData, error)
	/*
		FetchChannelInfo(channelId int64) (*ChannelInfo, error)
		FetchMediaDetail(mediaId int64) (*MediaDetail, error)
		FetchMediaFilter(channelId int64, pageSize, page int) (*MediaFilter, error)
		FetchMediaURL(mediaId, ci, source int64) (*MediaURL, error)
		FetchRecommendChannel(channelId int64) (*RecommendChannel, error)
	*/
}

type miVideoService struct {
	host   string
	api    string
	key    string
	token  string
	client *http.Client
}

func NewMiVideoService(host, api, key, token string) MiVideoService {
	return &miVideoService{
		host:   host,
		api:    api,
		key:    key,
		token:  token,
		client: &http.Client{},
	}
}

func (s *miVideoService) NewRequest(path string) *Request {
	return &Request{
		host:        s.host,
		api:         s.api,
		path:        path,
		key:         s.key,
		token:       s.token,
		paramValues: make(map[string]string),
		paramKeys:   make([]string, 0),
		client:      s.client,
	}
}

func (s *miVideoService) FetchHomeData() (data *types.HomeData, err error) {
	req := s.NewRequest("c/home")
	req.AddCommonParam()
	err = req.Get(func(r io.Reader) (err error) {
		json, err := simplejson.NewFromReader(r)
		if err != nil {
			return
		}

		data, err = types.ParseHomeData(json)
		if err != nil {
			return
			err = errors.ParseError(json)
		}
		return
	})

	return
}
