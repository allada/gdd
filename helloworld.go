package main
import "fmt"
func main() {
    helloWorldVar := "hello world"
    fmt.Println(helloWorldVar)
    helloWorldVar = "goodnight world"
    fmt.Println(helloWorldVar)
    PrintStuffOut()
    fmt.Println("hello world")
}

func PrintStuffOut() {
    fmt.Println("hello world")
    return
    fmt.Println("hello world")
}