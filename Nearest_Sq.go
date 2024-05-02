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

  if res == n {
    return n
  }
  if (res-n) < (n-res2) {
    return res
  }else {
    return res2
  }
}
