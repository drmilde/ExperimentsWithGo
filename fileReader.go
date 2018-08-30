package main

import (
       "fmt"
//       "math"
//      "runtime"
       )


func rec(p int) int {
     if p < 0 {
     	return 0
     } else {
       return p + (rec(p-1)) 
     }
}

func main() {
     ergebnis := rec(10000000)
     fmt.Println ("toll ", ergebnis)
}