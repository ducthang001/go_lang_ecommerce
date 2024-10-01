package basic

import (
	"testing"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2 // 3 -> fail
	)

	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	}

	// assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
	// assert.NotEqual(t, 2, 3)
	// assert.Nil(t, nil, nil)
}

func TestAddOne2(t *testing.T) {
	var (
		input  = 1
		output = 2 // 3 -> fail
	)

	actual := AddOne2(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	}

}

// func TestRequire(t *testing.T) {
// 	require.Equal(t, 2, 3) // block below code
// 	fmt.Println("Not executing")
// }

// func TestAssert(t *testing.T) {
// 	assert.Equal(t, 2, 3)
// 	fmt.Println("executing")
// }

// execute: go test ./... -coverprofile=cover
// execute: go tool cover -html=cover
