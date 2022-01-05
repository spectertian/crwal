package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Default struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

type DownStruct struct {
	Title string
	Url   string
	Type  string
}

type DownInfoStruct struct {
	CId         string `bson:"c_id"`
	Title       string
	Type        []string
	Url         string
	DownStatus  int          `bson:"down_status"`
	DownUrl     []DownStruct `bson:"down_url"`
	UpdatedTime time.Time    `bson:"updated_time"`
	CreatedTime time.Time    `bson:"created_time"`
}

type Dy struct {
	Url               string
	CId               string `bson:"c_id"`
	RId               string `bson:"r_id"`
	Title             string
	Alias             []string
	LongTitle         string `bson:"long_title"`
	Pic               string
	imgUrl            string `bson:"img_url"`
	Director          []string
	Stars             []string
	Introduction      string
	DownStatus        int    `bson:"down_status"`
	LastUpdated       string `bson:"last_updated"`
	UpdatedDate       string `bson:"updated_date"`
	Source            string
	UpdatedTime       time.Time `bson:"updated_time"`
	CreatedTime       time.Time `bson:"created_time"`
	ProductionDate    string    `bson:"production_date"`
	PageDate          string    `bson:"page_date"`
	Rating            string
	DoubanUrl         string `bson:"douban_url"`
	DoubanId          string `bson:"douban_id"`
	Tags              []string
	Type              []string
	Year              string
	Area              string
	RunTime           string `bson:"run_time"`
	Language          string
	DownUrl           []DownStruct `bson:"down_url"`
	ProductionCompany string       `bson:"production_company"`
	Status            string
	ClickCount        int `bson:"click_count"`
	DownCount         int `bson:"down_count"`
}
