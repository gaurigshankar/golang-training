package main


import (
  "fmt"
)

//In this example i am putting the entry in the channel created. Once the process of data
// ingestion is complete, i am specificcally closing hte channel

//Then ranging over channel and printing the content of the channel
// This iteration will happen till the close of the channel is done or the channel is empty

func main() {
  myDummyChannel := make(chan int)

  go func() {
    for i := 0; i <200; i++ {
      fmt.Println("before with ",i)
      myDummyChannel <- i
    }
    close(myDummyChannel)
  }()

  for n:= range myDummyChannel {
    fmt.Println(n)
  }
}
