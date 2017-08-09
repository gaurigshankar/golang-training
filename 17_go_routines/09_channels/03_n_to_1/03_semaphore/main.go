package main

import "fmt"


// Here in this example we are having one more channel that carries boolean to track
// if a given go routine is complete

// In each subroutine, once the for loop completes, we pass true to the ccmFlag channel
// In the third subroutine, we look for two for two values of ccmFlag channel , if we get two flags,
// then we close the channel

// the fmt.Println thats getting executed in the main thread, will be alive till the channel is closed 

func main() {
  myDummyChannel := make(chan int)
  myCCMSwitchChannel := make(chan bool)

  go func() {
    for i :=0; i<20; i++ {
      myDummyChannel <- i
    }
    myCCMSwitchChannel <- true
  }()
  go func() {
    for i :=0; i<20; i++ {
      myDummyChannel <- i
    }
    myCCMSwitchChannel <- true
  }()
  go func() {
    <- myCCMSwitchChannel
    <- myCCMSwitchChannel
    close(myDummyChannel)
  }()

  for n := range myDummyChannel {
    fmt.Println(n)
  }


}
