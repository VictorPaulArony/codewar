package main 
import (
        "fmt"
        )
        func main() {
            var rows int 
            //fmt.Scanln(&rows)
            rows = 9
            for i := 1; i<= rows; i++ {
                for sp := 1; sp <= rows - i; sp++{
                    fmt.Print(" ")
                }
                for j :=0; j != (2*i-1); j++ {
                    fmt.Printf("%d", i)
                    
                }
                
                fmt.Println()
                
            }
        }
