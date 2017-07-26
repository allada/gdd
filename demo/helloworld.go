package main
import "fmt"
func main() {
    helloWorldVar := "hello world"
    fmt.Println(helloWorldVar)
    helloWorldVar = "goodnight world"
    fmt.Println(helloWorldVar)
    waitChan := make(chan bool)
    go PrintStuffOut(waitChan)
    fmt.Println("hello world")
    <-waitChan
}

func PrintStuffOut(waitChan chan<- bool) {
    data := "hello world"
    fmt.Println(data)
    waitChan <-true
    return
    fmt.Println("hello world")
}