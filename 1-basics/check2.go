package main

import (
"fmt"
"runtime"
)

func main() {
// Functions can be used as expressions inside statements.
fmt. Println( runtime.NumCPU())
// You can use operators with function call expressions.
fmt. Println( runtime.NumCPU() + 1)
}
