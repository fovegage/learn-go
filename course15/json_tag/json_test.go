package json_tag

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type TestJson struct {
	Name string `json:"name" newtag:"newname"`
	Age  int    `json:"age" newtag:"newage"`
}

func Test1(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data := TestJson{}
	data.Name = "zhangsan"
	data.Age = 22
	byt, _ := json.Marshal(&data)
	fmt.Println(string(byt))

	var NewJson = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "newtag",
	}.Froze()

	byt, _ = NewJson.Marshal(&data)
	fmt.Println(string(byt))
}

func TestTag(t *testing.T) {
	// 判断不同的 标转化
	// {"name":"gage", "age":18}
	// {"newname":"gage", "newage":18}
	var user TestJson
	var NewJson = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "newtag",
	}.Froze()
	NewJson.Unmarshal([]byte(`{"newname":"gage", "newage":18}`), &user)
	println(user.Name, user.Age)

	data, _ := json.Marshal(user)
	println(string(data))

}
func Test2(t *testing.T) {
	user := User{
		Id:    0,
		Name:  "superlongstring",
		Bio:   "",
		Email: "foobar",
	}

	fmt.Println("Errors: ")
	for i, err := range validateStruct(user) {
		fmt.Printf("\t%d. %s\n", i+1, err.Error())
	}
}
