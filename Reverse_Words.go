// Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

// Examples
// "This is an example!" ==> "sihT si na !elpmaxe"
// "double  spaces"      ==> "elbuod  secaps" 
package main
import "strings"
func ReverseWords(str string) string {
  s := strings.Split(str, " ")

  for in, _ := range s{
    run := []rune(s[in])
    for i,j := 0 , len(run)-1;i < j ; i, j = i+1, j-1 {
      run[i], run[j] = run[j], run[i]
    }
    s[in] = string(run)
  }
  return strings.Join(s, " ")
  }
