package main

import("fmt"
"time"

) 

var c1 chan string
var c2 chan string


func customer(){ //customer


	fmt.Println(" prajwal i want mutton biryani in zomoto")

	for i :=0 ; i <2; i++{

		c1 <- "mutton biryani"

		fmt.Println("prajwal extra in cool drinks also")
	}
		fmt.Println("customer completed")
		
		time.Sleep(1)
	

}


func owner(){ //
	
   fmt.Println("welcome to zomoto")

   fmt.Println("what do u what sir")
	

	for i :=0 ; i <4; i++{

			//c2 <- "zomoto"
			fmt.Println("owner completed",<-c2)
	}
	
	
	time.Sleep(2)
	
}



func main (){

	c1 = make(chan string)
	c2 = make(chan string)

	fmt.Println("conservation b/w food ordering")

	go customer()
	go owner()


 	for i :=0; i <2; i++{

	fmt.Println("order is conform")
	select{
	case msg1 := <-c1:
		//c2 <-msg1
	fmt.Println("all items are good",msg1)
	//case msg2 := <-c2:
	//fmt.Println("have a nice day sir",msg2)

	
	}
}

}