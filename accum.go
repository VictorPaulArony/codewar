package main 
import (
  "fmt"
  "strinngs"
  )
func Accum(words string) string {
  // the space where to store the final result
  result := ""
  for i, word := range words {
    //change all the given chararcters into uppper case form first 
    up := strings.ToUpper(string(word))
    //itterate through all the strings up to i 
    rep := strings.Repeat(up, i)
    //change the itterated words to lower case 
    low := strings.ToLower(rep)
    //concatination of the lower case and the upper case 
    result += up + low 
    //adding of the - to all except to the n-1 word
    if i != len(words)-1 {
      result += "-" 
      }
    }
  return result
  }
   func main() {
	testCases := []string{"abcd", "RqaEzty", "cwAt"}
	for _, testCase := range testCases {
		fmt.Printf("accum(\"%s\") -> \"%s\"\n", testCase, Accum(testCase))
	}
}
