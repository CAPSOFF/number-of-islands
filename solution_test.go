package number_of_island

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var testCase = []struct {
	input [][]byte
	er    int
}{
	{
		[][]byte{
			[]byte("11000"),
			[]byte("01001"),
			[]byte("10011"),
			[]byte("00000"),
			[]byte("10110"),
		},
		5,
	},
	{
		[][]byte{
			[]byte("11110"),
			[]byte("11010"),
			[]byte("11000"),
			[]byte("00000"),
		},
		1,
	},
}

func Test_solution(t *testing.T) {
	assertion := assert.New(t)

	for _, tc := range testCase {
		fmt.Print("--------------------\n")
		for r, v1 := range tc.input {
			for c, _ := range v1 {
				p, _ := strconv.Atoi(string(tc.input[r][c]))
				fmt.Print(p, " ")

			}
			fmt.Print("\n")
		}

		result := solution(tc.input)
		fmt.Println("expected result : ", tc.er)
		fmt.Println("solution result : ", result)
		fmt.Print("--------------------\n\n")

		assertion.Equal(tc.er, result, "", tc)
	}
}
