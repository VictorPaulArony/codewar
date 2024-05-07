// The vowel substrings in the word codewarriors are o,e,a,io. The longest of these has a length of 2. Given a lowercase string that has alphabetic characters only (both vowels and consonants) and no spaces, return the length of the longest vowel substring. Vowels are any of aeiou.

// Good luck!

// If you like substring Katas, please try:
package main
import "strings"
func Solve(s string) int {
  var res int
  vol := "aeiou"
  count := 0
  for _, char := range s {
    if strings.Contains(vol, string(char)) {
      count ++
      if count > res {
        res = count
      }
    } else {
      count = 0
    }
   
  }
  
    return res
}
