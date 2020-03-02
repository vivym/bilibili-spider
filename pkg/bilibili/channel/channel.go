package channel

import (
	"net/http/cookiejar"
	"strconv"
	"time"

	"emperror.dev/errors"

	"github.com/go-resty/resty/v2"
)

var (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"
)

type Channel struct {
	http      *resty.Client
	channelID int
}

func New(cookieJar *cookiejar.Jar, channelID int) *Channel {
	http := resty.New().
		SetCookieJar(cookieJar).
		SetRetryCount(3).
		SetRetryWaitTime(5*time.Second).
		SetRetryMaxWaitTime(20*time.Second).
		SetHostURL("https://api.bilibili.com").
		SetHeader("User-Agent", userAgent)

	channel := Channel{
		http:      http,
		channelID: channelID,
	}
	return &channel
}

func (c *Channel) newlist(pageNum, pageSize int) (*NewListResult, error) {
	rsp, err := c.http.R().
		SetQueryParams(map[string]string{
			"rid":   strconv.Itoa(c.channelID),
			"type":  "0",
			"pn":    strconv.Itoa(pageNum),
			"ps":    strconv.Itoa(pageSize),
			"jsonp": "json",
			"_":     strconv.FormatInt(time.Now().UnixNano()/1000000, 10),
		}).
		SetResult(&NewListResult{}).
		Get("/x/web-interface/newlist")
	if err != nil {
		return nil, errors.WithMessage(err, "newlist http error")
	}

	result := rsp.Result().(*NewListResult)
	if result.Code != 0 {
		return result, errors.Errorf("newlist result: errorCode=%d errorMessage=%s", result.Code, result.Message)
	}
	return result, nil
}

func (c *Channel) archivesByTag(tagID, pageNum, pageSize int) (*NewListResult, error) {
	rsp, err := c.http.R().
		SetQueryParams(map[string]string{
			"tag_id": strconv.Itoa(tagID),
			"rid":    strconv.Itoa(c.channelID),
			"type":   "0",
			"pn":     strconv.Itoa(pageNum),
			"ps":     strconv.Itoa(pageSize),
			"jsonp":  "json",
			"_":      strconv.FormatInt(time.Now().UnixNano()/1000000, 10),
		}).
		SetResult(&NewListResult{}).
		Get("/x/tag/ranking/archives")
	if err != nil {
		return nil, errors.WithMessage(err, "archivesByTag http error")
	}

	result := rsp.Result().(*NewListResult)
	if result.Code != 0 {
		return result, errors.Errorf(
			"archivesByTag result: errorCode=%d errorMessage=%s",
			result.Code, result.Message,
		)
	}
	return result, nil
}

// order: click 播放数, scores 评论数, stow 收藏数, coin 硬币数, dm 弹幕数
// time format: 20200224/20200302
func (c *Channel) search(order string, pageNum, pageSize int, timeFrom, timeTo string, tagName string) (*SearchResult, error) {
	queryParams := map[string]string{
		"main_ver":    "v3",
		"search_type": "video",
		"view_type":   "hot_rank",
		"order":       order,
		"copy_right":  "-1",
		"cate_id":     strconv.Itoa(c.channelID),
		"page":        strconv.Itoa(pageNum),
		"pagesize":    strconv.Itoa(pageSize),
		"jsonp":       "json",
		"time_from":   timeFrom,
		"time_to":     timeTo,
		"_":           strconv.FormatInt(time.Now().UnixNano()/1000000, 10),
	}
	if tagName != "" {
		queryParams["keyword"] = tagName
	}
	rsp, err := c.http.R().
		SetQueryParams(queryParams).
		SetResult(&SearchResult{}).
		Get("https://s.search.bilibili.com/cate/search")
	if err != nil {
		return nil, errors.WithMessage(err, "search http error")
	}

	searchResult := rsp.Result().(*SearchResult)
	if searchResult.Code != 0 || searchResult.Msg != "success" {
		return searchResult, errors.New(searchResult.Msg)
	}
	return searchResult, nil
}
