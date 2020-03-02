package channel

type Query struct {
	channel *Channel

	tagID    int
	tagName  string
	timeFrom string
	timeTo   string
	order    string
	pageNum  int
	pageSize int

	numPages  int
	numVideos int

	result []*VideoInfo
	err    error
}

func (c *Channel) Q() *Query {
	query := Query{
		channel: c,

		order:    "time",
		pageNum:  1,
		pageSize: 20,
	}
	return &query
}

func (q *Query) SetTagID(tagID int) *Query {
	q.tagID = tagID
	return q
}

func (q *Query) SetTagName(tagName string) *Query {
	q.tagName = tagName
	return q
}

func (q *Query) SetTimeFrom(t string) *Query {
	q.timeFrom = t
	return q
}

func (q *Query) SetTimeTo(t string) *Query {
	q.timeTo = t
	return q
}

func (q *Query) SetTimeInterval(from, to string) *Query {
	q.timeFrom = from
	q.timeTo = to
	return q
}

// Order: time, click 播放数, scores 评论数, stow 收藏数, coin 硬币数, dm 弹幕数
func (q *Query) SetOrder(order string) *Query {
	q.order = order
	return q
}

func (q *Query) SetPageNum(pageNum int) *Query {
	q.pageNum = pageNum
	return q
}

func (q *Query) SetPageSize(pageSize int) *Query {
	q.pageSize = pageSize
	return q
}

func (q *Query) GetPageNum() int {
	return q.pageNum - 1
}

func (q *Query) GetNumPages() int {
	return q.numPages
}

func (q *Query) GetNumVideos() int {
	return q.numVideos
}

func (q *Query) GetResult() []*VideoInfo {
	return q.result
}

func (q *Query) GetError() error {
	return q.err
}

func (q *Query) NextPage() bool {
	defer func() { q.pageNum++ }()
	if q.order == "time" {
		var result *NewListResult
		var err error
		if q.tagID != 0 {
			result, err = q.channel.archivesByTag(q.tagID, q.pageNum, q.pageSize)
		} else {
			result, err = q.channel.newlist(q.pageNum, q.pageSize)
		}
		if err != nil {
			q.err = err
			return false
		}
		q.numVideos = result.Data.Page.Count
		q.numPages = int(float32(q.numVideos)/float32(q.pageSize) + 0.5)
		q.result = result.Data.Archives
		return len(q.result) > 0
	} else {
		result, err := q.channel.search(q.order, q.pageNum, q.pageSize, q.timeFrom, q.timeTo, q.tagName)
		if err != nil {
			q.err = err
			return false
		}
		q.numVideos = result.NumResults
		q.numPages = result.NumPages
		videoInfos := make([]*VideoInfo, 0, len(result.Result))
		for _, info := range result.Result {
			videoInfos = append(videoInfos, info.ToVideoInfo())
		}
		q.result = videoInfos
		return len(q.result) > 0
	}
}
