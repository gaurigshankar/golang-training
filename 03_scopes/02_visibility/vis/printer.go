package vis

import "fmt"

func GauriCustomPrinterFunction() {
  fmt.Printf("Inside GauriCustomPrinterFunction the value of MyName is %v \n",MyName
    // Though myTitle is not exported outside the package, they can be printed via exported functions as the function is still able to access within the package
  fmt.Printf("Inside GauriCustomPrinterFunction the value of myTitle is %v \n",myTitle)
}
