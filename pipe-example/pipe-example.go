package main

import (
	"github.com/yaronsumel/pipe"
	"fmt"
	"flag"
	"os"
	"time"
	"io"
)
// go get github.com/yaronsumel/pipe/example
// pipe-example --write | pipe-example
// small example to demonstrate code and pipe usage
func main() {
	var writeMode = flag.Bool("write", false, "write some demo data to stdout")
	flag.Parse()
	if *writeMode {
		writeData()
	}
	StdinChannel := make(pipe.StdDataChannel)
	go pipe.AsyncRead(pipe.Stdin, 1024, StdinChannel)
	for {
		select {
		case stdin := <-StdinChannel:
			if stdin.Err != nil {
				if stdin.Err == io.EOF {
					fmt.Printf("stdin-> EOF \r\n")
					os.Exit(1)
				}
				panic("stdin panic->" + stdin.Err.Error())
			}
			fmt.Printf("stdin-> `%s` \r\n", stdin.Data)
		}
	}
}

// writeData is a simple loop writing every second to stdout
func writeData() {
	for k := 0; k < 5; k++ {
		os.Stdout.Write([]byte("[example-writer] stdout:" + time.Now().String()))
		time.Sleep(time.Second * 1)
	}
	os.Exit(0)
}