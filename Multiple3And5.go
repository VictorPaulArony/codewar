package main

func Multiple3And5(num int) int {
  count := 0 
  for i := 1; i < num ; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      count += i
    }
  }
  return count
}
