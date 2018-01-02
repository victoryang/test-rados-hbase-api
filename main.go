package main

import (
"fmt"
"context"
"time"
"github.com/tsuna/gohbase"
"github.com/tsuna/gohbase/hrpc"
)

func testHbaseCase1 () {
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
}

func testHbaseCase2() {
	client := gohbase.NewClient("localhost")
	getRequest, err := hrpc.NewGetStr(context.Background(), "test", "rowkey1")
	if err != nil {
	}
	getRsp, err := client.Get(getRequest)
	if err !=nil {
	}
	fmt.Println(getRsp)
}

func main(){
	testHbaseCase1()
	time.Sleep(1)
	testHbaseCase2()
	fmt.Println("test finished, leaving...")
}
