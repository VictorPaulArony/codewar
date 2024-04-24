// Sort the Vowels!
// In this kata, we want to sort the vowels in a special format.

// Task
// Write a function which takes a input string s and return a string in the following way:

   
//                   C|                          R|
//                   |O                          n|
//                   D|                          d|
//    "CODEWARS" =>  |E       "Rnd Te5T"  =>      |
//                   W|                          T|
//                   |A                          |e
//                   R|                          5|
//                   S|                          T|

// Notes:

// List all the Vowels on the right side of |
// List every character except Vowels on the left side of |
// for the purpose of this kata, the vowels are : a e i o u
// Return every character in its original case
// Each line is seperated with \n
// Invalid input ( undefined / null / integer ) should return an empty string
package main
import ("strings"
        "fmt"
        )
func SortVowels(s string) string {
  if s == "" {
    return ""
  }
  res:= ""
  vol := "aeiouAEIOU"

  for _, word := range s {
    if  strings.Contains(vol, string(word)){
      res += fmt.Sprintf("|%s\n",string(word))
    } else {
      res += fmt.Sprintf("%s|\n",string(word))
      
    }
  }
	return res[:len(res)-1]
}
func main() {
	fmt.Println(SortVowels("CODEWARS"))
	fmt.Println(SortVowels("Rnd Te5T"))
}
