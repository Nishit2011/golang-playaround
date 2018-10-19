package main
 
import "fmt"
 
func main() {
   c := color("Red")
 
   fmt.Println(c.describe("is an awesome color"))
}
 
type color string
 
func (c color) describe(description string) (string) {
   return string(c) + " " + description
}
*****************************
explanation of above:
'describe' is a function with a receiver of type 'color' that requires an argument of type 'string', then returns a value of type 'string'


What is a byte slice?
A slice where every element inside corresponds to an ASCII character code.
It represents a string
data []byte ,is a byte slice

What is type conversion?
Takes value of one type and converts it into another type
for ex:
greeting := "Hi There!"
fmt.Println([]byte(greeting))

above will convert a greeting type string into byte slice


