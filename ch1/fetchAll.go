// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	ch := make(chan string)
// 	for _, url := range os.Args[1:] {
// 		if !strings.HasPrefix(url, "http://") {
// 			url = "http://" + url
// 		}
// 		go fetch(url, ch) // start a goroutine
// 	}
// 	for range os.Args[1:] {
// 		fmt.Println(<-ch) // receive from channel ch
// 	}
// 	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
// }

// func fetch(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err) // send to channel ch
// 		return
// 	}
// 	nbytes, err := io.Copy(io.Discard, resp.Body)
// 	resp.Body.Close() // don't leak resources
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err)
// 		return
// 	}
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

// }

package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)

func fetch(url string, file *os.File) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintln(file, fmt.Sprint(err))
        return
    }
    defer resp.Body.Close()

    nbytes, err := io.Copy(io.Discard, resp.Body)
    if err != nil {
        fmt.Fprintf(file, "while reading %s: %v\n", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    fmt.Fprintf(file, "%.2fs %7d %s\n", secs, nbytes, url)
}

func main() {
    url := "http://gopl.io"
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Run fetch twice without goroutines
    fetch(url, file)
    fetch(url, file)
}