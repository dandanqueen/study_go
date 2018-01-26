// dup version 2
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    files := os.Args[1:]
    for _, file := range files {
        countLines(file)
    }
}

func countLines(fname string) {
    counts := make(map[string]int)

    f, err := os.Open(fname)
    if err != nil {
        fmt.Fprintf(os.Stderr, "practice4: %v\n", err)
        return
    }
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
        if counts[input.Text()] > 1 {
            fmt.Println(fname)
            break
        }
    }
    f.Close()
}
