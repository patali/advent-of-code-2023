package utils

import (
	"fmt"
	"time"
)

func PrintTimeElapsed(inStart time.Time, inName string) {
	elapsed := time.Since(inStart)
	fmt.Printf("%s took %s\n", inName, elapsed)
	fmt.Println("-----")
}
