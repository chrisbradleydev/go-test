package unpack

import (
	"fmt"
)

func ExecUnpack() {
    arr := []int{1,3,5,7,9}
    sum := unpack(arr...)
    fmt.Printf("Sum %d\n", sum)
}

func unpack(args ...int) int {
   sum := 0
   for _, v := range args {
      sum = sum + v
   }
   return sum
}