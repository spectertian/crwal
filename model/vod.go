package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type WVods struct {
	VodId            string
	TypeId           string
	TypeIdParent     string
	GroupId          string
	VodName          string
	VodSub           string
	VodEn            string
	VodStatus        string
	VodLetter        string
	VodColor         string
	VodTag           string
	VodClass         string
	VodPic           string
	PicID            string
	vodPicThumb      string
	vodPicSlide      string
	vodPicScreenshot string
	vodActor         string
	vodDirector      string
	vodWriter        string
	vodBehind        string
	vodBlurb         string
	vodRemarks       string
	vodPubdate       string
	VodTotal         string
	VodSerial        string
	VodTv            string
	VodWeekday       string
	VodArea          string
	VodLang          string
	VodYear          string
	VodVersion       string
	VodState         string
	VodAuthor        string
	VodJumpurl       string
	VodTpl           string
	VodTplPlay       string
	VodTplDown       string
	VodIsSend        string
	VodLock          string
	VodLevel         string
	VodCopyright     string
	VodPoints        string
	VodPointsPlay    string
	VodPointsDown    string
	VodHits          string
	VodHitsDay       string
	VodHitsWeek      string
	VodHitsMonth     string
	VodDuration      string
	VodUp            string
	VodDown          string
	VodScore         string
	VodScoreAll      string
	VodScoreNum      string
	VodTime          string
	VodTimeAdd       string
	VodTimeHits      string
	VodTimeMake      string
	VodTrysee        string
	VodDoubanId      string
	VodDoubanScore   string
	VodReUrl         string
	VodReVod         string
	VodReArt         string
	VodPwd           string
	VodPwdUrl        string
	VodPwdPlay       string
	VodPwdPlayUrl    string
	VodPwdDown       string
	VodPwdDownUrl    string
	VodContent       string
	VodPlayFrom      string //"wjyun$$$wjm3u8",
	VodPlayServer    string //"no$$$no",
	VodPlayNote      string //$$$
	VodPlayUrl       string //"DVD$https://v4.cdtlas.com/share/KJhnO2ZC23aiA8NO$$$DVD$https://v4.cdtlas.com/20220212/P4WVmj2q/index.m3u8",
	VodDownFrom      string
	VodDownServer    string
	VodDownNote      string
	VodDownUrl       string
	VodPlot          string
	VodPlotName      string
	VodPlotDetail    string
	TypeName         string
}

type WJVod struct {
	VodId            int       `json:"vod_id"`
	TypeId           int       `json:"type_id"`
	TypeIdParent     int       `json:"type_id_1"`
	GroupId          int       `json:"group_id"`
	VodName          string    `json:"vod_name"`
	VodSub           string    `json:"vod_sub"`
	VodEn            string    `json:"vod_en"`
	VodStatus        int       `json:"vod_status"`
	VodLetter        string    `json:"vod_letter"`
	VodColor         string    `json:"vod_color"`
	VodTag           string    `json:"vod_tag"`
	VodClass         string    `json:"vod_class"`
	VodPic           string    `json:"vod_pic"`
	vodPicThumb      string    `json:"vod_pic_thumb"`
	vodPicSlide      string    `json:"vod_pic_slide"`
	vodPicScreenshot string    `json:"vod_pic_screenshot"`
	vodActor         string    `json:"vod_actor"`
	vodDirector      string    `json:"vod_director"`
	vodWriter        string    `json:"vod_writer"`
	vodBehind        string    `json:"vod_behind"`
	vodBlurb         string    `json:"vod_blurb"`
	vodRemarks       string    `json:"vod_remarks"`
	vodPubdate       string    `json:"vod_pubdate"`
	VodTotal         int       `json:"vod_total"`
	VodSerial        string    `json:"vod_serial"`
	VodTv            string    `json:"vod_tv"`
	VodWeekday       string    `json:"vod_weekday"`
	VodArea          string    `json:"vod_area"`
	VodLang          string    `json:"vod_lang"`
	VodYear          string    `json:"vod_year"`
	VodVersion       string    `json:"vod_version"`
	VodState         string    `json:"vod_state"`
	VodAuthor        string    `json:"vod_author"`
	VodJumpurl       string    `json:"vod_jumpurl"`
	VodTpl           string    `json:"vod_tpl"`
	VodTplPlay       string    `json:"vod_tpl_play"`
	VodTplDown       string    `json:"vod_tpl_down"`
	VodIsSend        int       `json:"vod_isend"`
	VodLock          int       `json:"vod_lock"`
	VodLevel         int       `json:"vod_level"`
	VodCopyright     int       `json:"vod_copyright"`
	VodPoints        int       `json:"vod_points"`
	VodPointsPlay    int       `json:"vod_points_play"`
	VodPointsDown    int       `json:"vod_points_down"`
	VodHits          int       `json:"vod_hits"`
	VodHitsDay       int       `json:"vod_hits_day"`
	VodHitsWeek      int       `json:"vod_hits_week"`
	VodHitsMonth     int       `json:"vod_hits_month"`
	VodDuration      string    `json:"vod_duration"`
	VodUp            int       `json:"vod_up"`
	VodDown          int       `json:"vod_down"`
	VodScore         string    `json:"vod_score"`
	VodScoreAll      int       `json:"vod_score_all"`
	VodScoreNum      int       `json:"vod_score_num"`
	VodTime          string    `json:"vod_time"`
	VodTimeAdd       int       `json:"vod_time_add"`
	VodTimeHits      int       `json:"vod_time_hits"`
	VodTimeMake      int       `json:"vod_time_make"`
	VodTrysee        int       `json:"vod_trysee"`
	VodDoubanId      int       `json:"vod_douban_id"`
	VodDoubanScore   string    `json:"vod_douban_score"`
	VodReUrl         string    `json:"vod_reurl"`
	VodReVod         string    `json:"vod_rel_vod"`
	VodReArt         string    `json:"vod_rel_art"`
	VodPwd           string    `json:"vod_pwd"`
	VodPwdUrl        string    `json:"vod_pwd_url"`
	VodPwdPlay       string    `json:"vod_pwd_play"`
	VodPwdPlayUrl    string    `json:"vod_pwd_play_url"`
	VodPwdDown       string    `json:"vod_pwd_down"`
	VodPwdDownUrl    string    `json:"vod_pwd_down_url"`
	VodContent       string    `json:"vod_content"`
	VodPlayFrom      string    `json:"vod_play_from"`
	VodPlayServer    string    `json:"vod_play_server"`
	VodPlayNote      string    `json:"vod_play_note"`
	VodPlayUrl       string    `json:"vod_play_url"`
	VodDownFrom      string    `json:"vod_down_from"`
	VodDownServer    string    `json:"vod_down_server"`
	VodDownNote      string    `json:"vod_down_note"`
	VodDownUrl       string    `json:"vod_down_url"`
	VodPlot          int       `json:"vod_plot"`
	VodPlotName      string    `json:"vod_plot_name"`
	VodPlotDetail    string    `json:"vod_plot_detail"`
	TypeName         string    `json:"type_name"`
	ImgUrl           string    `json:"img_url"`
	UpdatedTime      time.Time `bson:"updated_time"`
	CreatedTime      time.Time `bson:"created_time"`
}
type JsonResult struct {
	Code      int     `json:"code"`
	Msg       string  `json:"msg"`
	Page      int     `json:"page"`
	PageCount int     `json:"pagecount"`
	Limit     string  `json:"limit"`
	Total     int     `json:"total"`
	List      []WJVod `json:"list"`
}

type VodIndexHas struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	VodDoubanId int                `json:"vod_douban_id"`
}
