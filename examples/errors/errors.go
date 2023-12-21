package main

import (
	"fmt"
	"errors"
)

func FunctionWithError(val int) (int, error) {
	if val == -1 {
		return -1, errors.New("NOT A VALID VALUE")
	}

	return val, nil
}

// Custom error defintion
/*
It’s possible to use custom types as errors by implementing 
the Error() method on them.
*/

type UDFError struct {
	args int
	prob string
}

func (udferr *UDFError) Error() string {
	return fmt.Sprintf("%d - %s", udferr.args, udferr.prob)
}

func FuncWithUDFError(val int) (int, error) {
	if val == 1 {
		return 1, &UDFError{args: 1, prob: "INVALID"}
	}
	return -1, nil
}

func main() {	
	res, err := FunctionWithError(-1)
	fmt.Println(res, err)

	r, e := FuncWithUDFError(1)
	fmt.Println(r, e)
}