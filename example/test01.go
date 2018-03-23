package main

import (
	"fmt"
	simpleJson "github.com/gdaisukesuzuki/go-simplejson"
	"encoding/json"
)


func main() {
  var jsonStr = `{"Name": "Ken", "Age": 32}`
  var jsonStr2 = `{"Name": "Taro", "Age": 4}`
  var jsonStr3 = `{"Name": "Hanako", "Age": 2, "Sex":"Female"}`

  js, err := simpleJson.NewJson([]byte(jsonStr))
  if err != nil {
    panic(err)
  }
  
  var fbyte2 interface{}
  err2 := json.Unmarshal([]byte(jsonStr2), &fbyte2)
  if err2 != nil {
    panic(err2)
  }
  var fbyte3 interface{}
  err3 := json.Unmarshal([]byte(jsonStr3), &fbyte3)
  if err3 != nil {
    panic(err3)
  }
  
   js.ZeroIndex("Children")
   js.AddIndex("Children",0,fbyte2)
   k,_ := js.EncodePretty()
  fmt.Println(string(k)) // {Ken, 32}
  js.Get("Children").GetIndex(0).Set("Sex","Male")
   k,_ = js.EncodePretty()
  fmt.Println(string(k)) // {Ken, 32}
   js.AddIndex("Children",-1,fbyte3)
   k,_ = js.EncodePretty()
  fmt.Println(string(k)) // {Ken, 32}

}

