package channel

import (
	"strconv"
	"time"
)

type NewListResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Archives []*VideoInfo `json:"archives"`
		Page     PageInfo     `json:"page"`
	} `json:"data"`
}

type PageInfo struct {
	Count int `json:"count"`
	Num   int `json:"num"`
	Size  int `json:"size"`
}

type VideoInfo struct {
	AvID      int    `json:"aid" bson:"aid"`
	Videos    int    `json:"videos" bson:"videos"`
	TagID     int    `json:"tid" bson:"tid"`
	TagName   string `json:"tname" bson:"tname"`
	Copyright int    `json:"copyright" bson:"copyright"`
	Pic       string `json:"pic" bson:"pic"`
	Title     string `json:"title" bson:"title"`
	PubDate   int    `json:"pubdate" bson:"pubdate"`
	Ctime     int    `json:"ctime" bson:"ctime"`
	Desc      string `json:"desc" bson:"desc"`
	State     int    `json:"state" bson:"state"`
	Attribute int    `json:"attribute" bson:"attribute"`
	Duration  int    `json:"duration" bson:"duration"`
	MissionID int    `json:"mission_id" bson:"mission_id"`
	Rights    struct {
		BP            int `json:"bp" bson:"bp"`
		Elec          int `json:"elec" bson:"elec"`
		Download      int `json:"download" bson:"download"`
		Movie         int `json:"movie" bson:"movie"`
		Pay           int `json:"pay" bson:"pay"`
		HD5           int `json:"hd5" bson:"hd5"`
		NoReprint     int `json:"no_reprint" bson:"no_reprint"`
		Autoplay      int `json:"autoplay" bson:"autoplay"`
		UGCPay        int `json:"ugc_pay" bson:"ugc_pay"`
		IsCooperation int `json:"is_cooperation" bson:"is_cooperation"`
		UGCPayPreview int `json:"ugc_pay_preview" bson:"ugc_pay_preview"`
		NoBackground  int `json:"no_background" bson:"no_background"`
	} `json:"rights" bson:"rights"`
	Owner   OwnerInfo `json:"owner" bson:"owner"`
	Stat    StatInfo  `json:"stat" bson:"stat"`
	Dynamic string    `json:"dynamic" bson:"dynamic"`
	CID     int       `json:"cid" bson:"cid"`
	BVID    string    `json:"bvid" bson:"bvid"`
}

type OwnerInfo struct {
	MID  int    `json:"mid" bson:"mid"`
	Name string `json:"name" bson:"name"`
	Face string `json:"face" bson:"face"`
}

type StatInfo struct {
	AvID     int `json:"aid" bson:"aid"`
	View     int `json:"view" bson:"view"`
	Danmaku  int `json:"danmaku" bson:"danmaku"`
	Reply    int `json:"reply" bson:"reply"`
	Favorite int `json:"favorite" bson:"favorite"`
	Coin     int `json:"coin" bson:"coin"`
	Share    int `json:"share" bson:"share"`
	NowRank  int `json:"now_rank" bson:"now_rank"`
	HisRank  int `json:"his_rank" bson:"his_rank"`
	Like     int `json:"like" bson:"like"`
	Dislike  int `json:"dislike" bson:"dislike"`
}

type SearchResult struct {
	Code           int      `json:"code"`
	ShowModuleList []string `json:"show_module_list"`
	CostTime       struct {
		ParamsCheck         string `json:"params_check"`
		IllegalHandler      string `json:"illegal_handler"`
		AsResponseFormat    string `json:"as_response_format"`
		AsRequest           string `json:"as_request"`
		DeserializeResponse string `json:"deserialize_response"`
		AsRequestFormat     string `json:"as_request_format"`
		Total               string `json:"total"`
		MainHandler         string `json:"main_handler"`
	} `json:"cost_time"`
	Result         []*SearchVideoInfo `json:"result"`
	ShowColumn     int                `json:"show_column"`
	RqtType        string             `json:"rqt_type"`
	NumPages       int                `json:"numPages"`
	NumResults     int                `json:"numResults"`
	CrrQuery       string             `json:"crr_query"`
	PageSize       int                `json:"pagesize"`
	SuggestKeyword string             `json:"suggest_keyword"`
	Cache          int                `json:"cache"`
	ExpBits        int                `json:"exp_bits"`
	ExpStr         string             `json:"exp_str"`
	SeID           string             `json:"seid"`
	Msg            string             `json:"msg"`
	EggHit         int                `json:"egg_hit"`
	Page           int                `json:"page"`
}

type SearchVideoInfo struct {
	SendDate     int    `json:"senddate"`
	RankOffset   int    `json:"rank_offset"`
	Tag          string `json:"tag"`
	Duration     int    `json:"duration"`
	ID           int    `json:"id"`
	RankScore    int    `json:"rank_score"`
	BadgePay     bool   `json:"badgepay"`
	PubDate      string `json:"pubdate"`
	Author       string `json:"author"`
	Review       int    `json:"review"`
	MID          int    `json:"mid"`
	IsUnionVideo int    `json:"is_union_video"`
	RankIndex    int    `json:"rank_index"`
	Type         string `json:"type"`
	ArcRank      string `json:"arcrank"`
	Play         string `json:"play"`
	Pic          string `json:"pic"`
	Description  string `json:"description"`
	VideoReview  int    `json:"video_review"`
	IsPay        int    `json:"is_pay"`
	Favorites    int    `json:"favorites"`
	ArcURL       string `json:"arcurl"`
	BVID         string `json:"bvid"`
	Title        string `json:"title"`
}

func (s *SearchVideoInfo) ToVideoInfo() *VideoInfo {
	viewCnt, _ := strconv.Atoi(s.Play)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	pubDateObj, _ := time.ParseInLocation("2006-01-02 15:04:05", s.PubDate, loc)
	pubDate := int(pubDateObj.Unix())
	return &VideoInfo{
		AvID:     s.ID,
		TagName:  s.Tag,
		Pic:      s.Pic,
		Title:    s.Title,
		PubDate:  pubDate,
		Ctime:    pubDate,
		Desc:     s.Description,
		Duration: s.Duration,
		Owner: OwnerInfo{
			MID:  s.MID,
			Name: s.Author,
		},
		Stat: StatInfo{
			AvID:     s.ID,
			View:     viewCnt,
			Danmaku:  s.VideoReview,
			Reply:    s.Review,
			Favorite: s.Favorites,
		},
		BVID: s.BVID,
	}
}
