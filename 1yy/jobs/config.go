package main

import "github.com/robfig/config"

type JobsConfig struct {
	MiHost   string
	MiApi    string
	MiKey    string
	MiToken  string
	MySQLDsn string
}

func NewJobsConfig(confPath string) (jobsConf *JobsConfig, err error) {
	conf, err := config.ReadDefault("./jobs.conf")
	if err != nil {
		return
	}

	host, err := conf.String("mivideo", "host")
	if err != nil {
		return
	}

	api, err := conf.String("mivideo", "api")
	if err != nil {
		return
	}

	key, err := conf.String("mivideo", "key")
	if err != nil {
		return
	}

	token, err := conf.String("mivideo", "token")
	if err != nil {
		return
	}

	dsn, err := conf.String("mysql", "dsn")
	if err != nil {
		return
	}

	jobsConf = &JobsConfig{
		MiHost:   host,
		MiApi:    api,
		MiKey:    key,
		MiToken:  token,
		MySQLDsn: dsn,
	}
	return
}
