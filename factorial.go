package main
import "fmt"

func factorial(num int, ans int) int {
	if num <= 1{
		return ans
}
	ans*=num
	return factorial(num-1,ans)
}

func main(){
	var num int 
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	res := factorial(num,1)
	fmt.Printf("\nThe factorial of %d is %d",num, res)
}
