// Package mock provides ...
package mock

import "fmt"

func Printf(format string, a ...interface{}) {
	fmt.Printf("mock: "+format, a...)
}
