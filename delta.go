package main 

import (
	
	"fmt"
	"math"
)

func closez(x float64) float64 {

	var k,a,b,z,y,i float64
    
    k = math.Sqrt(x)
	for z = 1; z < 10; z++{
        for i = 1; i < 10; i++{

            a = z*z - x
            b = a / 2*z
			z = z - b
            
            y= z-k
            fmt.Println(y)
            

		}
        
	 	
	 	
	 	
	}
    
    return y
}

	
	

	



func main() {
	closez(4)
	
}