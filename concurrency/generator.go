// Generator: the function that returns a channel
// Taken from - https://youtu.be/f6kdp27TYZs?t=866 (Go Concurrency Patterns Rob Pike)
// Run on - https://go.dev/play/p/RdN1_DR913o
package main

func boring(msg string) <-chan string {
  c := make(chan string)
  go func() {
    for i:=0; ;i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
    }
  }()
  return c // return the channel to the caller
} 

func main() {
  c := boring("boring!")
  for i:=0; i<5; i++ {
    fmt.Printf("You say %s\n", <-c)
  }
  fmt.Println("You are boring I'm leaving")
}


