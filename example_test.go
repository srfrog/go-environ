// Copyright 2014-present Codehack. All rights reserved.
// For mobile and web development visit http://codehack.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package environ_test

import (
	"fmt"
	"github.com/codehack/go-environ"
)

var testEnv = environ.NewEnv()

func ExamplePut() {
	testEnv.Put("first", "value")
	testEnv.Put("second", "value")
	testEnv.Put("third", "value")
	testEnv.Print()
	// Output:
	// 0 = first=value
	// 1 = second=value
	// 2 = third=value
}

func ExamplePut_change() {
	testEnv.Put("first", "true")
	testEnv.Put("second", "something")
	testEnv.Put("third", "-981")
	testEnv.Put("PI", "3.141592653589793")
	testEnv.Print()
	// Output:
	// 0 = first=true
	// 1 = second=something
	// 2 = third=-981
	// 3 = PI=3.141592653589793
}

func ExampleIndex() {
	fmt.Println(testEnv.Index("first"))
	fmt.Println(testEnv.Index("unknown"))
	fmt.Println(testEnv.Index("third"))
	// Output:
	// 0
	// -1
	// 2
}

func ExampleContains() {
	fmt.Println(testEnv.Contains("first"))
	fmt.Println(testEnv.Contains("second"))
	fmt.Println(testEnv.Contains("fourth"))
	// Output:
	// true
	// true
	// false
}

func ExampleGet() {
	fmt.Println(testEnv.Get("first"))
	fmt.Println(testEnv.Get("wrong"))
	fmt.Println(testEnv.Get("right"))
	// Output:
	// true
}

func ExampleGetBool() {
	fmt.Println(testEnv.GetBool("first"))
	fmt.Println(testEnv.GetBool("second"))
	fmt.Println(testEnv.GetBool("third"))
	fmt.Println(testEnv.GetBool("PI"))
	// Output:
	// true
	// false
	// false
	// false
}

func ExampleGetFloat() {
	fmt.Println(testEnv.GetFloat("first"))
	fmt.Println(testEnv.GetFloat("second"))
	fmt.Println(testEnv.GetFloat("third"))
	fmt.Println(testEnv.GetFloat("PI"))
	// Output:
	// 0
	// 0
	// -981
	// 3.141592653589793
}

func ExampleGetInt() {
	fmt.Println(testEnv.GetInt("first"))
	fmt.Println(testEnv.GetInt("second"))
	fmt.Println(testEnv.GetInt("third"))
	fmt.Println(testEnv.GetInt("PI"))
	// Output:
	// 0
	// 0
	// -981
	// 0
}

type whatever struct{}

func ExampleSet() {
	testEnv.Set("first", 1.01)
	err := testEnv.Set("second", whatever{})
	if err != nil {
		fmt.Println(err.Error())
	}
	testEnv.Set("third", 1.0)
	testEnv.Print()
	// Output:
	// environ.Set: second: value "environ_test.whatever" is not supported
	// 0 = first=1.01
	// 1 = second=something
	// 2 = third=1
	// 3 = PI=3.141592653589793
}

// Testing file environ_test.go has plenty of tests to
// serve as examples too. Enjoy.
