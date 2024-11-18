package variables

import "fmt"

func PrintVariables() {
	var name string = "Akhilesh Chandewar (string)"
	fmt.Println(name)
	var age int = 24
	fmt.Println("Age is " , age , "(int)")
	var height float64 = 5.7
	fmt.Println("Height is " , height , " feets (float32 / float64)")
	var graduated bool = true
	fmt.Println("Graduated" , graduated , "(boolean)")
}