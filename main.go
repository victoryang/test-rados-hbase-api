package main

import (
"fmt"
"context"
"github.com/tsuna/gohbase"
"github.com/tsuna/gohbase/hrpc"
)

func main(){
client := gohbase.NewClient("localhost")
values := map[string]map[string][]byte{"cf1": map[string][]byte{"col1": []byte{0}}}
putRequest, err := hrpc.NewPutStr(context.Background(), "test", "rowkey1", values)
if err != nil {
fmt.Println(err)
}
rsp, err := client.Put(putRequest)
if err != nil {
fmt.Println(rsp)
fmt.Println(err)
}
client.Close()
fmt.Println("hello world")
}
