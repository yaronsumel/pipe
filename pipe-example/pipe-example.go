package main

import (
	"flag"
	"fmt"
	"github.com/yaronsumel/pipe"
	"io"
	"os"
	"time"
)

// go get github.com/yaronsumel/pipe/example
// pipe-example --write | pipe-example
// small example to demonstrate code and pipe usage
func main() {
	var writeMode = flag.Bool("write", false, "write some demo data to stdout")
	flag.Parse()
	if *writeMode {
		writeData()
		return
	}
	StdinChannel := make(pipe.StdDataChannel)
	go pipe.AsyncRead(pipe.Stdin, 1024*128, StdinChannel)
	for {
		select {
		case stdin := <-StdinChannel:
			if stdin.Err != nil {
				fmt.Println(stdin.Err)
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
		os.Stdout.Write([]byte(""))
		time.Sleep(time.Second * 1)
	}
	os.Exit(1)
}
