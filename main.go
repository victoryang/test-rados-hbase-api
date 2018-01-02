package main

import (
"fmt"
"github.com/tsuna/gohbase"
)

func main(){
client := gohbase.NewClient("localhost")
values := map[string]map[string][]byte{"cf1": map[string][]byte{"col1": []byte{aaa}}}
putRequest, err := hrpc.NewPutStr(context.Background(), "test", "rowkey1", values)
if err != nil {
}
fmt.Println("hello world")
}
