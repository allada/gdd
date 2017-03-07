package main
import "fmt"
func main() {
    fmt.Println("hello world")
    PrintStuffOut()
    fmt.Println("hello world")
}

func PrintStuffOut() {
    fmt.Println("hello world")
    return
    fmt.Println("hello world")
}