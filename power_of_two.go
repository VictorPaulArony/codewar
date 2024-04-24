// Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).

package main

func PowersOfTwo(n int) []uint64 {
  var res []uint64
  for i:= 0 ; i <= n; i++ {
    p:= uint64(1 << i)
    res = append(res, p)
  }
  return res
}
