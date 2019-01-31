package watcher

import "fmt"

// Main wraps Run, adding a log.Fatal(err) on error, and setting the log formatter
func Main() {
	fmt.Println("Watcher")
}
