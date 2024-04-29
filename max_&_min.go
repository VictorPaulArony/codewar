finding the maximun and minimum values in a given array
package main 
import "fmt"

func MinMax(arr []int) (int,int) {
    min := arr[0]
    max := arr[0]
    for _, num := range arr {
        if num < min {
            min = num 
        }
        if num >max {
            max = num
        }
    }
    return min, max
}
func main() {
    arr := []int{3, 5, 10, 9, 2, 7}
    min, max := MinMax(arr)
    fmt.Printf("Min: %d, Max: %d\n", min, max)
}
