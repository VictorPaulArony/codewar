/*You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z. Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).

Example:
a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13
Bash note:
input : 2 strings with substrings separated by ,
output: number as a string*/
package main

func MxDifLg(a1 []string, a2 []string) int {
    // your code\
  if len(a1)== 0 || len(a2) == 0 {
    return -1
  }
  var res1 int 
  var res2 int
  min1 := len(a1[0])
  max1 := len(a1[0])
  max2 := len(a2[0])
  min2 := len(a2[0])
  for _, word := range a1 {
    l1 := len(word)
    if l1 < min1 {
      min1 = l1
    }
    if l1 > max1 {
      max1 = l1
    } 
  }
  for _, word := range a2 {
    l2 := len(word)
    if l2 > max2 {
      max2 = l2
    }
    if l2 < min2 {
      min2 = l2
    }
  }

  res1 = max1 - min2
  res2 = max2 - min1
  if res1 > res2 {
    return res1
  }
  return res2
}
