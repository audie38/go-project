package main

import "fmt"

func logging() {
	message := recover()
	if message != nil{
		fmt.Println("Error: ", message)
	}else{
		fmt.Println("Finish Executing Function")
	}
}

func runApplication(value int){
	defer logging()
	result := 10/value;
	fmt.Println("Run Application", result)
}

func runApp(err bool){
	defer logging()
	if(err){
		panic("TEST ERROR")
	}
	fmt.Println("App Running")
}

func main() {
	//defer -> can be scheduled to run after a certain function already executed, it will be executed even if previous function errors
	//panic -> similar to throw exception
	runApp(false)
	runApp(true)
	runApplication(1)
}