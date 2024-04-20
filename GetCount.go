// Return the number (count) of vowels in the given string.

// We will consider a, e, i, o, u as vowels for this Kata (but not y).

// The input string will only consist of lower case letters and/or spaces.
package main 
import "fmt"
func myGetCount(str string) (count int) {
  count = 0
  vowels := string("aeiou")
  for _, s := range str {
    for _, vowel := range vowels {
      if s == vowel {
        count++
        }
      }
    }
  return count
  }
func main() {
	// Test cases
	fmt.Println(countVowels("hello"))      // Output: 2
	fmt.Println(countVowels("world"))      // Output: 1
	fmt.Println(countVowels("programming")) // Output: 3
}
