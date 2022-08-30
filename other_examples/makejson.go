package main
import(
  "encoding/json"
  "fmt"

  // "strings"
)


func main(){
  m := make(map[string]string)
  var name string
  var address string
  fmt.Printf("\n Enter the Name: ")
  fmt.Scanf("%s",&name)
  fmt.Printf("\nEnter address : ")
  fmt.Scanf("%s",&address)
  m["name"] = name
  m["address"]=address
  fmt.Println("\nmap :",m)
  barr,err := json.Marshal(m)

  if(err==nil){
    mapstring := string(barr)
    fmt.Println("\njson Object  :", mapstring)
  }
}
