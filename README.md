# pipe [![Go Report Card](https://goreportcard.com/badge/github.com/yaronsumel/pipe)](https://goreportcard.com/report/github.com/yaronsumel/pipe)  [![GoDoc](https://godoc.org/github.com/yaronsumel/pipe?status.svg)](https://godoc.org/github.com/yaronsumel/pipe)
small package for reading pipe data

 Installation
```bash
go get -u github.com/yaronsumel/pipe
```
      
 Sync Usage
Read is Sync Action to get all pipe data fits in the predifened size
```go
	data,err := pipe.Read(pipe.Stdin,1024)
	if err!=nil{
		//do something with the error
	}
  ```
  
      
 Async Usage
AyncRead will keep reading from the pipe and write it back to StdDataChannel. Don't forget to handle the errors.
```go
	StdinChannel := make(pipe.StdDataChannel)
	go pipe.AsyncRead(pipe.Stdin, 1024, StdinChannel)
	for {
		select {
		case stdin := <-StdinChannel:
			if stdin.Err != nil {
				panic(stdin.Err)
			}
			fmt.Println(stdin.Data)
		}
	}
  ```
  
 Example
```bash
go get -u github.com/yaronsumel/pipe/pipe-example 
pipe-example --write | pipe-example
```
      