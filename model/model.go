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
type Dy struct {
	Url               string
	CId               string `bson:"c_id"`
	RId               string `bson:"r_id"`
	Title             string
	Alias             []string
	LongTitle         string `bson:"long_title"`
	Pic               string
	Director          []string
	Stars             []string
	Introduction      string
	DownStatus        int
	LastUpdated       string `bson:"last_updated"`
	UpdatedDate       string `bson:"updated_date"`
	Source            string
	UpdateTime        time.Time `bson:"update_time"`
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
