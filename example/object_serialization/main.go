package main

import (
	"fmt"
	"github.com/chaksunshine/go_helper/store"
	"github.com/chaksunshine/go_helper/store/enum"
	"time"
)

func main() {

	var ur = store.OSStruct.Find(enum.RegisterTypeFile, "ur", new(userRecord)).(*userRecord).init() //  new(userRecord).init()
	fmt.Println(ur)
	time.Sleep(time.Hour)
}
