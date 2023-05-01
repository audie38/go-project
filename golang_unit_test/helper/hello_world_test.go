package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

- Skip() : to skip unit test

- stretchr : assert, require
- assert functions : if unit test failed -> will trigger Fail()
- require functions : if unit test failed -> will trigger FailNow()
- functions: Equal, NotEqual, Nil, NotNil
- params: (t, [expected result], [actual result], [message])
*/

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Milson")
	assert.Equal(t, "Hello Audie", result, "Result must be Hello Audie")
}

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

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("Unit Test Cannot run on Windows OS")
	}

	result := HelloWorld("Audie")
	require.Equal(t, "Hello Audie", result)
}