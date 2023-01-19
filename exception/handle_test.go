package exception

import (
	"errors"
	"fmt"
	"testing"
)

func TestWarp(t *testing.T) {

	err1 := errors.New("error1")

	strErr := fmt.Errorf("return err %w", StrError{"error1"})

	innerErr := fmt.Errorf("inner error is %w", err1)

	fmt.Printf("strErr: %v, innerErr: %v \r\n", strErr, innerErr)

	// codeErrEg := &CodeError{}
	// fmt.Println("Error.as", errors.As(codeErr, &codeErrEg))
	fmt.Println("Error.as", errors.As(strErr, &StrError{}))
	fmt.Println("Error.is", errors.Is(strErr, err1))
	// fmt.Println("Error.as", errors.As(innerErr, &codeErrEg))
	fmt.Println("Error.as", errors.As(innerErr, &StrError{}))
	fmt.Println("Error.is", errors.Is(innerErr, err1))
}
