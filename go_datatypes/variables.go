package main; 

import "fmt"

func main(){
	//There are two ways to declare variables 

	//var and const 
	//var can be changed once it declared 

	var x float64 = 10 ;
	const y = 20; 

	//Short declaration 
	x += 1;
	// y += 1;


	//Third way to assign a variable quickly; 
	z := 30; //you don't have to specify type or valye 

	fmt.Println("Result: ", x);
	fmt.Println("Result: ", y);
	fmt.Println("Result: ", z);

	//Variable declaration without initial values 

	// var a string ; when no initial values, then it will be empty string "";
	// var b int; When no initial values, then it will be 0(int), for float (0)
	// var c bool; when no initial values, then it will be false(0)
	var a string ; 
	var b int; 
	var c float32; 
	var d bool;

	a = "Welcome to Go Code!"
	b = 10; 
	d = true; 

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Println(d);

	//Difference between var and := 
	/*var - can be used inside and outside of function 
			Variable declaration and value assignment can be done separately 

	  := - can only be used inside functions 
	  	   variables declarations and value assignment cannot be done separately(must be done in the same line)
		   
	*/
	

}