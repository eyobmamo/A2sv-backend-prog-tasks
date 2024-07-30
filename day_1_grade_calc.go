// the code calculate student grade 
package main

import "fmt"

func calcAverage(result map[string]float32) float32 {
	var total float32
	var average float32
	
	leng := float32(len(result))
	for _,v := range result{
		total += float32(v)
	}
	average = total/leng
	return average

}

func main(){

	var name string
	var no_courses int

	var result = make(map[string]float32)

	fmt.Println("grade calculator")
	fmt.Println("Enter Your Name:")
	fmt.Scanln(&name)
	fmt.Println("enter no of subject you take :")
	fmt.Scanln(&no_courses)

	
    
	for i:=0; i<no_courses; i++{
		var subject string
		var score float32

		fmt.Println("enter subject name :")
		fmt.Scanln(&subject)
		
		
		for {
			fmt.Println("enter score b/n 0 and 100:")
			fmt.Scanln(&score)
			if score >= 0 && score <=100{
				break
			}else{
				fmt.Println("Invalid score! ")
			}	
			}
		result[subject]=score	
		}
		

	fmt.Printf(" \n\nName :%v \n",name)
	fmt.Printf("Number of Subject teken:%v \n\n",no_courses)
	for k,v := range result {
		fmt.Printf("%v : %v \n" ,k,v)
	}

	// calculate the average score using calcAverange
	avg := calcAverage(result)
	fmt.Printf("Your average score is: %.2f",avg)
	}
