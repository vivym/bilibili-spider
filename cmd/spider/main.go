package main

import (
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/vivym/bilibili-spider/pkg/bilibili/channel"

	"github.com/Kamva/mgm/v2"
	"github.com/vivym/bilibili-spider/internal/db"

	"github.com/vivym/bilibili-spider/pkg/bilibili"
)

type VideoInfo struct {
	mgm.DefaultModel  `bson:",inline"`
	channel.VideoInfo `bson:",inline"`
}

func main() {
	db.SetupDB("mongodb://localhost:27017/", "bilibili")

	b := bilibili.New()
	channel := b.NewChannel(124)

	numCreated := 0
	query := channel.Q()
	for query.NextPage() {
		fmt.Printf("page #%d: %d\n", query.GetPageNum(), numCreated)
		for _, videoInfo := range query.GetResult() {
			model := &VideoInfo{mgm.DefaultModel{}, *videoInfo}
			if err := mgm.Coll(model).Create(model); err != nil {
				if merr, ok := err.(mongo.WriteException); ok {
					for _, werr := range merr.WriteErrors {
						if werr.Code != 11000 {
							fmt.Println("insert video error:", werr)
						}
					}
				} else {
					fmt.Println("insert video unknown error:", err)
				}
			} else {
				numCreated++
			}
		}

		delay := time.Duration(200 + rand.Intn(200))
		time.Sleep(delay * time.Millisecond)
	}
	if err := query.GetError(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("done.", numCreated)
}
