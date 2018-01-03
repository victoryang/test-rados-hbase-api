package main

import (
	"fmt"
	"context"
	"time"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

// Insert a cell
func testHbaseCase1 () {
	client := gohbase.NewClient("localhost")
	values := map[string]map[string][]byte{"cf1": map[string][]byte{"col1": []byte{0}}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "test", "rowkey1", values)
	if err != nil {
		//fmt.Println(err)
	}
	rsp, err := client.Put(putRequest)
	if err != nil {
		fmt.Println(rsp)
		//fmt.Println(err)
	}
	client.Close()
}

// Get an entire row
func testHbaseCase2() {
	client := gohbase.NewClient("localhost")
	getRequest, err := hrpc.NewGetStr(context.Background(), "test", "rowkey1")
	if err != nil {
	}
	getRsp, err := client.Get(getRequest)
	if err !=nil {
	}
	//fmt.Println(getRsp)
	for _, cell := range getRsp.Cells {
		fmt.Println(string(cell.Row))
		fmt.Println(string(cell.Family))
		fmt.Println(string(cell.Qualifier))
		fmt.Println(string(cell.Value))
	}
	client.Close()
}

// Get a specific cell
func testHbaseCase3() {
	client := gohbase.NewClient("localhost")
	family := map[string][]string{"cf1": []string{"col1"}}
	getRequest, err := hrpc.NewGetStr(context.Background(), "test", "rowkey1", hrpc.Families(family))
	if err != nil {
	}
	getRsp, err := client.Get(getRequest)
	if err != nil {
	}
	for _, cell := range getRsp.Cells {
		fmt.Println(string(cell.Row))
		fmt.Println(string(cell.Family))
		fmt.Println(string(cell.Qualifier))
		fmt.Println(string(cell.Value))
	}
	client.Close()
}

func main(){
	testHbaseCase1()
	time.Sleep(1)
	testHbaseCase2()
	time.Sleep(1)
	testHbaseCase3()
	fmt.Println("test finished, leaving...")
}
