package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, j, x, y, modus int
	y = 0
	fmt.Scan(&n)
	for i = 0; i <n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		for j = i+1; j < n; j++{
			if A[i] == A[j]{
				x++
			}
		}
		if y < x{
			y = x
			modus = A[i]
		}
	}
	fmt.Print(modus)
}