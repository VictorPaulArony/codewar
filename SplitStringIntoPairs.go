// Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').

// Examples:

// * 'abc' =>  ['ab', 'c_']
// * 'abcdef' => ['ab', 'cd', 'ef']
package main 
import "fmt"
func Solution(str string) []string {
  var res []string 
  for i := 0;i < len(str);i +=2 {
    par := str[i:i+2]
    if len(str)%2 != 0{
      str += "_" 
      }
    res = append(res, par)
    }
  return res
  }
    func main() {
	input := ".abcdefghd"
	pairs := Solution(input)
	fmt.Println(pairs)
}
      
