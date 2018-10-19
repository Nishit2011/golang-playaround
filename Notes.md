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


How does random number is generated in GoLang?

newPosition := rand.Intn(len(d) - 1)
        rand.Intn is a sudo random generator
		go by default uses a random number generator that depends on
		some seed value
		seed value is passed in the generator and then the generator is used to make a random sequence
		of numbers or values or whatever we're in randomising
		so go by default always uses the exact same seed. So every single time we restart our program as soon as we
		start trying to generate random numbers we're always going to get the exact same sequence because we always start from the exact seed value





