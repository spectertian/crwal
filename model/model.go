package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Default struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	DownStatus     int                `bson:"down_status"`
	LongTitle      string             `bson:"long_title"`
	Pic            string
	imgUrl         string `bson:"img_url"`
	Director       []string
	Stars          []string
	Area           string
	Rating         string
	Introduction   string
	ProductionDate string `bson:"production_date"`
}

type UpdateHas struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	DownStatus int                `bson:"down_status"`
}

type IndexHas struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

type DefaultTopicStruct struct {
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
	LongTitle   string `bson:"long_title"`
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
	Status            int
	ClickCount        int `bson:"click_count"`
	DownCount         int `bson:"down_count"`
}

type UpDyStruct struct {
	LongTitle   string    `bson:"long_title"`
	UpdatedDate string    `bson:"updated_date"`
	UpdatedTime time.Time `bson:"updated_time"`
	DoubanUrl   string    `bson:"douban_url"`
	Year        string
	Area        string
	RunTime     string `bson:"run_time"`
	Language    string
}

type UpdateDyStruct struct {
	LongTitle      string       `bson:"long_title"`
	DownUrl        []DownStruct `bson:"down_url"`
	DownStatus     int          `bson:"down_status"`
	DoubanUrl      string       `bson:"douban_url"`
	Rating         string
	ProductionDate string    `bson:"production_date"`
	UpdatedTime    time.Time `bson:"updated_time"`
}

type FDyStruct struct {
	ID     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Url    string
	CId    string `bson:"c_id"`
	RId    string `bson:"r_id"`
	Title  string
	Pic    string
	imgUrl string `bson:"img_url"`
}

type Update struct {
	Url            string
	InfoId         string `bson:"info_id"`
	CId            string `bson:"c_id"`
	Title          string
	Date           string
	Type           string
	ProductionDate string    `bson:"production_date"`
	CreatedTime    time.Time `bson:"created_time"`
	UpdatedTime    time.Time `bson:"updated_time"`
}

type UpUpdateStruct struct {
	Title          string
	Date           string
	ProductionDate string    `bson:"production_date"`
	UpdatedTime    time.Time `bson:"updated_time"`
}

type NewsStruct struct {
	Url            string
	InfoId         string `bson:"info_id"`
	CId            string `bson:"c_id"`
	Title          string
	Date           string
	ProductionDate string    `bson:"production_date"`
	CreatedTime    time.Time `bson:"created_time"`
	UpdatedTime    time.Time `bson:"updated_time"`
}

type UpNewsStruct struct {
	Title          string
	Date           string
	ProductionDate string    `bson:"production_date"`
	UpdatedTime    time.Time `bson:"updated_time"`
}

type TopicStruct struct {
	Url         string
	NId         int    `bson:"n_id"`
	FilmNum     string `bson:"film_num"`
	Title       string
	Content     string
	Date        string
	CreatedTime time.Time `bson:"created_time"`
}

type TopicListStruct struct {
	Url          string
	NId          int    `bson:"n_id"`
	InfoId       string `bson:"info_id"`
	TopicId      string `bson:"topic_id"`
	CId          string `bson:"c_id"`
	Introduction string
	Title        string
	Director     []string
	Stars        []string
	Rating       string
	imgUrl       string `bson:"img_url"`
	Pic          string
	Area         string
	CreatedTime  time.Time `bson:"created_time"`
}

type UpdateIndexListStruct struct {
	Title          string
	Date           string
	ProductionDate string    `bson:"production_date"`
	UpdatedTime    time.Time `bson:"updated_time"`
}
type IndexListStruct struct {
	Url            string
	Type           string
	Sort           int
	InfoId         string `bson:"info_id"`
	CId            string `bson:"c_id"`
	Title          string
	Date           string
	ProductionDate string    `bson:"production_date"`
	CreatedTime    time.Time `bson:"created_time"`
	UpdatedTime    time.Time `bson:"updated_time"`
}

type BZYStruct struct {
	TPId              string `bson:"t_p_id"`
	TPTitle           string `bson:"t_p_title"`
	TId               string `bson:"t_id"`
	TTitle            string `bson:"t_title"`
	Url               string
	CId               string `bson:"c_id"`
	Title             string
	EmTitle           string `bson:"em_title"`
	Alias             []string
	LongTitle         string `bson:"long_title"`
	Pic               string
	imgUrl            string `bson:"img_url"`
	Director          []string
	Stars             []string
	Introduction      string
	DownStatus        int       `bson:"down_status"`
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
	Status            int
	ClickCount        int `bson:"click_count"`
	DownCount         int `bson:"down_count"`
}

type TPlayStruct struct {
	Title string
	Url   string
}

type TKLStruct struct {
	Title string
	List  []TPlayStruct
}

type TKStruct struct {
	TPId              string `bson:"t_p_id"`
	TPTitle           string `bson:"t_p_title"`
	TId               string `bson:"t_id"`
	TTitle            string `bson:"t_title"`
	Url               string
	CId               string `bson:"c_id"`
	Title             string
	EmTitle           string `bson:"em_title"`
	Alias             []string
	LongTitle         string `bson:"long_title"`
	Pic               string
	imgUrl            string `bson:"img_url"`
	Director          []string
	Stars             []string
	Introduction      string
	DownStatus        int       `bson:"down_status"`
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
	Play              []TKLStruct
	RunTime           string `bson:"run_time"`
	Language          string
	DownUrl           []DownStruct `bson:"down_url"`
	ProductionCompany string       `bson:"production_company"`
	Status            int
	ClickCount        int `bson:"click_count"`
	DownCount         int `bson:"down_count"`
}

type TKUpdateStruct struct {
	EmTitle     string    `bson:"em_title"`
	UpdatedTime time.Time `bson:"updated_time"`
	DoubanId    string    `bson:"douban_id"`
	Play        []TKLStruct
}

type TKUpdateIntroductionStruct struct {
	Introduction string
}
