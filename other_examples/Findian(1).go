package main
import ("fmt"
		"strings"
		)

func findIan (testText string) bool {
testText = strings.ToLower(testText)
if strings.Contains(testText,"a"){
	if strings.HasPrefix(testText,"i"){
		if strings.HasSuffix(testText,"n"){
		
		return true
		}
		}
	}  
	return false
}
	
func main() {

var testText string 
fmt.Printf("enter a text that begins with 'i', contains 'a' and ends with 'n'  :   ")
	fmt.Scan(&testText,0)

for findIan(testText)==false {
	fmt.Printf("not good enough... try again: enter a text that begins with 'i', contains 'a' and ends with 'n'  :   ")
	fmt.Scan(&testText,0)
	findIan(testText)
}
fmt.Printf("we found Ian!!")
}



