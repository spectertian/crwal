package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Wiki struct {
	WikiId        int `bson:"wiki_id"`
	Title         string
	PostImage     string `bson:"post_image"`
	ImageUrl      string `bson:"img_url"`
	Rating        float64
	Alias         []string
	Tags          []string
	Year          int
	Area          []string
	Director      []string
	Stars         []string
	Writes        []string
	Episodes      string
	Language      []string
	EpisodesTime  string `bson:"episodes_time"`
	RunTime       string `bson:"run_time"`
	FirstPlayDate string `bson:"first_play_date"`
	IMDb          string `bson:"imdb"`
	Introduction  string
	UpdatedTime   time.Time `bson:"updated_time"`
	CreatedTime   time.Time `bson:"created_time"`
}

type WikiIndexHas struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	WikiId      int                `json:"wiki_id"`
	CreatedTime time.Time          `bson:"created_time"`
}
