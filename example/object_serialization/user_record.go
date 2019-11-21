package main

import (
	"fmt"
	"github.com/chaksunshine/go_helper/store"
	"github.com/chaksunshine/go_helper/store/dao"
	"github.com/chaksunshine/go_helper/store/enum"
	"time"
)

type userRecord struct {
	Idx     int
	Content []int
}

func (this *userRecord) init() *userRecord {
	if len(this.Content) == 0 {
		this.Content = make([]int, 0)
	}

	store.OSStruct.Register(enum.RegisterTypeFile, dao.Register{
		Object: this,
		Name:   "ur",
		Second: time.Second * 5,
	})

	go this.start()
	return this
}

func (this *userRecord) start() {
	for range time.NewTicker(time.Millisecond * 90).C {
		this.Idx++
		this.Content = append(this.Content, this.Idx)
		fmt.Println(this.Idx)

		if this.Idx >= 300 {
			this.Idx = 0
			this.Content = []int{}
		}
	}
}
