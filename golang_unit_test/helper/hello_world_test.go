package helper

import (
	"fmt"
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
- go test -run [func]/[sub test name] : run Specific SubTest

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

func TestMain(m *testing.M){
	// Before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
	// After
}

func TestSubTest(t *testing.T){
	t.Run("Audie", func(t *testing.T){
		result := HelloWorld("Audie")
		require.Equal(t, "Hello Audie", result, "Result must be Hello Audie")
	})

	t.Run("Milson", func(t *testing.T){
		result := HelloWorld("Milson")
		require.Equal(t, "Hello Milson", result, "Result must be Hello Milson")
	})
}

type Tests struct{
	name string
	request string
	expected string
	message string
}

func TestHelloWorldTable(t *testing.T){
	tests := []Tests{
		{
			name: "HelloWorld(Ichigo)",
			request: "Ichigo",
			expected: "Hello Ichigo",
			message: "Result must be Hello Ichigo",
		},
		{
			name: "HelloWorld(Kurosaki)",
			request: "Kurosaki",
			expected: "Hello Kurosaki",
			message: "Result must be Hello Kurosaki",
		},
	}

	for i:= 0; i < len(tests); i++{
		t.Run(tests[i].name, func(t *testing.T){
			result := HelloWorld(tests[i].request)
			assert.Equal(t, tests[i].expected, result, tests[i].message)
		})
	}
}