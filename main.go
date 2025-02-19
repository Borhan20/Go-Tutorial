package main 

import "fmt"
import "main/prime"

func main(){
	var primes []int 
	primes = prime.GetPrime(50)
	fmt.Println(primes)
	fmt.Println("hello world")
}