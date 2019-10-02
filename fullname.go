package main
 
import "fmt"
 
func main() {


        fmt.Print("Enter Firstname String: ")  
        var Firstname string    
        fmt.Scanln(&Firstname)                 
    fmt.Print("Enter Lastname String: ")
    var Lastname string
    fmt.Scanln(&Lastname)
        fmt.Print(Firstname + Lastname)             
}