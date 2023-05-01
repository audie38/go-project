package helper

import "testing"

/**
Notes:
- go test : run all test function
- go test -v : run test and show what func tested
- go test -v -run [func name] : run specific test function
- go test ./... : run go test from root folder

- Fail() : failed the test but keep continue
- FailNow() : failed the unit test and stop executing code below within the fucntion

- Error() : to add unit test failed message, keep continue
- Fatal() : to add unit test failed message, but unit test stop executing code below within the fucntion
*/

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Audie")
	if result != "Hello Audie"{
		// t.Fail()
		t.Error("Result Not Hello Audie")
	}
}

func TestHelloWorld2(t *testing.T){
	result := HelloWorld("Audie")
	if result != "Hello Audie"{
		// t.FailNow()
		t.Fatal("Result must be Hi Audie")
	}
}