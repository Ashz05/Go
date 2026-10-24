package main; 

import ("fmt")

func main(){

	//Variable name with more than one word can be diffcult to read 
	//Camel Case,Pascal Case,Snake Case 

	// var a, b, c, d int = 1,3,5,7.4;
	// fmt.Println(a);
	// fmt.Println(b);
	// fmt.Println(c);
	// fmt.Println(d);
	
	//Multiple declaration in One Line 
	//if we provide an type in same line then we can't declare different type of variables in the same line 

	var name , age , address  = "Jhon", 21, "Clover Street 2nd Lane";
	fmt.Println(name,age,address);

	//Mutltiple variable declarations can also be grouped together into 
	//a block for greater readability

	var(
		a int ;
		b int = 1;
		c string = "Hello";
	)

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	
}