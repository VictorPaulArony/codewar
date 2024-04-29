// Write reverseList function that simply reverses lists.
package main 
func ReverseList(list []int) []int {
    for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
        list[i], list[j] = list[j], list[i]
    }
    return list
}
