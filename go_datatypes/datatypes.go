package main; 
//In go, There are different types of a variables 

/*
int = stores integers(whole numbers), such as 123 or -123 
float32 - stores floating numbers, with devcimal, such as 19.99 or -19.99 
string - stores text such as "Hello World"
bool - stores values with two states : true or false 
*/
import "fmt"

func main(){
	//String - single or double quotes, sequence of characters enclosed in single quotes or double quotes 
	//int - integer value, -ve to +ve values 
	//float32 or float 64 , decimal point

	var name string = "Jhon";
	var age int = 23; 
	var is_alive bool = true; 

	fmt.Println(name,age,is_alive);


}