package bilibili

import (
	"net/http/cookiejar"

	"github.com/vivym/bilibili-spider/pkg/bilibili/channel"

	"golang.org/x/net/publicsuffix"
)

type Bilibili struct {
	cookieJar *cookiejar.Jar
}

func New() *Bilibili {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	b := Bilibili{
		cookieJar: cookieJar,
	}
	return &b
}

func (b *Bilibili) NewChannel(channelID int) *channel.Channel {
	return channel.New(b.cookieJar, channelID)
}
