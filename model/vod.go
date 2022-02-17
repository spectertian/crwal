package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type WJVod struct {
	VodId            int       `bson:"vod_id" json:"vod_id"`
	TypeId           int       `bson:"type_id" json:"type_id"`
	TypeIdParent     int       `bson:"type_id_1" json:"type_id_1"`
	GroupId          int       `bson:"group_id" json:"group_id"`
	VodName          string    `bson:"vod_name" json:"vod_name"`
	VodSub           string    `bson:"vod_sub" json:"vod_sub"`
	VodEn            string    `bson:"vod_en" json:"vod_en"`
	VodStatus        int       `bson:"vod_status" json:"vod_status"`
	VodLetter        string    `bson:"vod_letter" json:"vod_letter"`
	VodColor         string    `bson:"vod_color" json:"vod_color"`
	VodTag           string    `bson:"vod_tag" json:"vod_tag"`
	VodClass         string    `bson:"vod_class" json:"vod_class"`
	VodPic           string    `bson:"vod_pic" json:"vod_pic"`
	VodPicThumb      string    `bson:"vod_pic_thumb" json:"vod_pic_thumb"`
	VodPicSlide      string    `bson:"vod_pic_slide" json:"vod_pic_slide"`
	VodPicScreenshot string    `bson:"vod_pic_screenshot" json:"vod_pic_screenshot"`
	VodActor         string    `bson:"vod_actor" json:"vod_actor"`
	VodDirector      string    `bson:"vod_director" json:"vod_director"`
	VodWriter        string    `bson:"vod_writer" json:"vod_writer"`
	VodBehind        string    `bson:"vod_behind" json:"vod_behind"`
	VodBlurb         string    `bson:"vod_blurb" json:"vod_blurb"`
	VodRemarks       string    `bson:"vod_remarks" json:"vod_remarks"`
	VodPubdate       string    `bson:"vod_pubdate" json:"vod_pubdate"`
	VodTotal         int       `bson:"vod_total" json:"vod_total"`
	VodSerial        string    `bson:"vod_serial" json:"vod_serial"`
	VodTv            string    `bson:"vod_tv" json:"vod_tv"`
	VodWeekday       string    `bson:"vod_weekday" json:"vod_weekday"`
	VodArea          string    `bson:"vod_area" json:"vod_area"`
	VodLang          string    `bson:"vod_lang" json:"vod_lang"`
	VodYear          string    `bson:"vod_year" json:"vod_year"`
	VodVersion       string    `bson:"vod_version" json:"vod_version"`
	VodState         string    `bson:"vod_state" json:"vod_state"`
	VodAuthor        string    `bson:"vod_author" json:"vod_author"`
	VodJumpurl       string    `bson:"vod_jumpurl" json:"vod_jumpurl"`
	VodTpl           string    `bson:"vod_tpl" json:"vod_tpl"`
	VodTplPlay       string    `bson:"vod_tpl_play" json:"vod_tpl_play"`
	VodTplDown       string    `bson:"vod_tpl_down" json:"vod_tpl_down"`
	VodIsSend        int       `bson:"vod_isend" json:"vod_isend"`
	VodLock          int       `bson:"vod_lock" json:"vod_lock"`
	VodLevel         int       `bson:"vod_level" json:"vod_level"`
	VodCopyright     int       `bson:"vod_copyright" json:"vod_copyright"`
	VodPoints        int       `bson:"vod_points" json:"vod_points"`
	VodPointsPlay    int       `bson:"vod_points_play" json:"vod_points_play"`
	VodPointsDown    int       `bson:"vod_points_down" json:"vod_points_down"`
	VodHits          int       `bson:"vod_hits" json:"vod_hits"`
	VodHitsDay       int       `bson:"vod_hits_day" json:"vod_hits_day"`
	VodHitsWeek      int       `bson:"vod_hits_week" json:"vod_hits_week"`
	VodHitsMonth     int       `bson:"vod_hits_month" json:"vod_hits_month"`
	VodDuration      string    `bson:"vod_duration" json:"vod_duration"`
	VodUp            int       `bson:"vod_up" json:"vod_up"`
	VodDown          int       `bson:"vod_down" json:"vod_down"`
	VodScore         string    `bson:"vod_score" json:"vod_score"`
	VodScoreAll      int       `bson:"vod_score_all" json:"vod_score_all"`
	VodScoreNum      int       `bson:"vod_score_num" json:"vod_score_num"`
	VodTime          string    `bson:"vod_time" json:"vod_time"`
	VodTimeAdd       int       `bson:"vod_time_add" json:"vod_time_add"`
	VodTimeHits      int       `bson:"vod_time_hits" json:"vod_time_hits"`
	VodTimeMake      int       `bson:"vod_time_make" json:"vod_time_make"`
	VodTrysee        int       `bson:"vod_trysee" json:"vod_trysee"`
	VodDoubanId      int       `bson:"vod_douban_id" json:"vod_douban_id"`
	VodDoubanScore   string    `bson:"vod_douban_score" json:"vod_douban_score"`
	VodReUrl         string    `bson:"vod_reurl" json:"vod_reurl"`
	VodReVod         string    `bson:"vod_rel_vod" json:"vod_rel_vod"`
	VodReArt         string    `bson:"vod_rel_art" json:"vod_rel_art"`
	VodPwd           string    `bson:"vod_pwd" json:"vod_pwd"`
	VodPwdUrl        string    `bson:"vod_pwd_url" json:"vod_pwd_url"`
	VodPwdPlay       string    `bson:"vod_pwd_play" json:"vod_pwd_play"`
	VodPwdPlayUrl    string    `bson:"vod_pwd_play_url" json:"vod_pwd_play_url"`
	VodPwdDown       string    `bson:"vod_pwd_down" json:"vod_pwd_down"`
	VodPwdDownUrl    string    `bson:"vod_pwd_down_url" json:"vod_pwd_down_url"`
	VodContent       string    `bson:"vod_content" json:"vod_content"`
	VodPlayFrom      string    `bson:"vod_play_from" json:"vod_play_from"`
	VodPlayServer    string    `bson:"vod_play_server" json:"vod_play_server"`
	VodPlayNote      string    `bson:"vod_play_note" json:"vod_play_note"`
	VodPlayUrl       string    `bson:"vod_play_url" json:"vod_play_url"`
	VodDownFrom      string    `bson:"vod_down_from" json:"vod_down_from"`
	VodDownServer    string    `bson:"vod_down_server" json:"vod_down_server"`
	VodDownNote      string    `bson:"vod_down_note" json:"vod_down_note"`
	VodDownUrl       string    `bson:"vod_down_url" json:"vod_down_url"`
	VodPlot          int       `bson:"vod_plot" json:"vod_plot"`
	VodPlotName      string    `bson:"vod_plot_name" json:"vod_plot_name"`
	VodPlotDetail    string    `bson:"vod_plot_detail" json:"vod_plot_detail"`
	TypeName         string    `bson:"type_name" json:"type_name"`
	ImgUrl           string    `bson:"img_url" json:"img_url""`
	UpdatedTime      time.Time `bson:"updated_time"`
	CreatedTime      time.Time `bson:"created_time"`
}
type JsonResult struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Page      interface{} `json:"page"`
	PageCount int         `json:"pagecount"`
	Limit     string      `json:"limit"`
	Total     int         `json:"total"`
	List      []WJVod     `json:"list"`
}

type VodIndexHas struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	VodDoubanId int                `json:"vod_douban_id"`
	CreatedTime time.Time          `bson:"created_time"`
}
