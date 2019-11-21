package str

import (
	"encoding/json"
	"fmt"
	"os"
)

// 打印json
func PrintJson(val interface{}) {

	bytes, e := json.Marshal(val)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(bytes))
	os.Exit(2)
}
