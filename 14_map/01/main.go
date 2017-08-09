package main

import "fmt"

func main() {
  sampleMap := make(map[string] int)
  sampleMap["gauri"] = 1
  sampleMap["shankar"] = 2

  fmt.Println(sampleMap)

  value, isPresent := sampleMap["gauri"]
  fmt.Println("Existing Key", value, isPresent)

  value, isPresent = sampleMap["dummyKey"]
  fmt.Println("NonExisting Key ", value, isPresent)


  anotherMap()
  yetAnotherMap()
  deleteFromMap()
  lengthFromMap()
}


func anotherMap() {
  anotherMap := make(map[string] string)

  fmt.Println("---------------------------------")
  anotherMap["fName"] = "Gauri"
  anotherMap["lName"] = "Shankar"
  anotherMap["place"] = "Antartica"

  value, isPresent := anotherMap["fName"]
  fmt.Println("First Name ", value, isPresent)

    value = anotherMap["place"]
      fmt.Println("Place ", value)
  fmt.Println("---------------------------------")
}

func yetAnotherMap() {
  yetAnotherMap := map[string] string {"fName" : "gauri", "lName" : "shankar" , "place" :
  "Antartica"}

  fmt.Println("---------------------------------")
  value, isPresent := yetAnotherMap["fName"]
  fmt.Println("First Name ", value, isPresent)

    value = yetAnotherMap["place"]
      fmt.Println("Place ", value)

  fmt.Println("---------------------------------")
}

func deleteFromMap(){
  yetAnotherMap := map[string] string {"fName" : "gauri", "lName" : "shankar" , "place" :
  "Antartica"}
  fmt.Println("---------------------------------")
    fmt.Println(yetAnotherMap)
  delete(yetAnotherMap, "fName")
  fmt.Println(yetAnotherMap)
  yetAnotherMap["language"] = "English"
  fmt.Println(yetAnotherMap)
  fmt.Println("---------------------------------")
}


func lengthFromMap(){
  yetAnotherMap := map[string] string {"fName" : "gauri", "lName" : "shankar" , "place" :
  "Antartica"}
  fmt.Println("---------------------------------")
    fmt.Println(len(yetAnotherMap))
    delete(yetAnotherMap, "fName")
    fmt.Println(len(yetAnotherMap))
    yetAnotherMap["language"] = "English"
    fmt.Println(len(yetAnotherMap))
    fmt.Println("---------------------------------")
}
