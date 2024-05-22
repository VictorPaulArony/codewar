package main 
func Number(busStops [][2]int) int {
  stoper := 0
  for _, stop := range busStops {
    stoper += stop[0] - stop[1] 
    }
  return stoper
  }
// func main() {
// 	// Example usage:
// 	busStops := [][2]int{
// 		{3, 0}, // 3 people get on, 0 get off
// 		{4, 2}, // 4 people get on, 2 get off
// 		{2, 1}, // 2 people get on, 1 gets off
// 	}

// 	result := Number(busStops)
// 	fmt.Printf("Remaining passengers on the bus: %d\n", result)
// }
