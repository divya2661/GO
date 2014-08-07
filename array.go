package main 

import "fmt"
var pow = []int{1, 2, 3, 4, 5, 16, 27, 12, 134, }
func main() {
	
	for i, v := range pow{

		fmt.Printf("2**%d=%d\n",i,v)
	}


	p := []int{2, 3, 4, 5, 7, 11, 12}
	fmt.Println("p ==",p);

	a := make([]int, 5)
	printSlice("a",a)

	b := make([]int,0,5)
	printSlice("b",b)

	c := b[:2]
	printSlice("c",c)

	d := c[2:5]
	printSlice("d",d)


	for i:=0; i<len(p); i++{

		fmt.Printf("p[%d] == %d\n",i,p[i])
}
	fmt.Println("P[:3]==", p[:3])
	fmt.Println("p[:4]==", p[4:])


}

func printSlice(s string,x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
	
}