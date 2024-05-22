/*Your task is to find the nearest square number, nearest_sq(n) or nearestSq(n), of a positive integer n.

For example, if n = 111, then nearest\_sq(n) (nearestSq(n)) equals 121, since 111 is closer to 121, the square of 11, than 100, the square of 10.

If the n is already the perfect square (e.g. n = 144, n = 81, etc.), you need to just return n.

Good luck :)

Check my other katas:*/

package main

func NearestSq(n int) int {
  num := 1
    for num*num < n {
    num ++
  }
  res := num * num
  res2 := (num-1) * (num-1)
  //Return the number (count) of vowels in the given string.

  // We will consider a, e, i, o, u as vowels for this Kata (but not y).
  
  // The input string will only consist of lower case letters and/or spaces.

  if res == n {
    return n
  }
  if (res-n) < (n-res2) {
    return res
  }else {
    return res2
  }
}
// package kata

// func NearestSq(n int) int {
//   var min, max int
  
//   for i := 1; i < n*n; i++ {
    
//     if i*i == n {
//       return n
//     }
    
//     if i*i < n {
//       min = i*i
//     }
    
//     if i*i > n {
//       max = i*i
//       if n - min < max - n {
//         return min
//       }
//       return max
//     }
//   }
//   return n
// }
