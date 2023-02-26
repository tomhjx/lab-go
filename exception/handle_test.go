package exception

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, errors.As(strErr, &StrError{}), true)
	assert.Equal(t, errors.Is(strErr, err1), false)

	// fmt.Println("Error.as", errors.As(innerErr, &codeErrEg))
	fmt.Println("Error.as", errors.As(innerErr, &StrError{})) //false
	fmt.Println("Error.is", errors.Is(innerErr, err1))        // true
	assert.Equal(t, errors.As(innerErr, &StrError{}), false)
	assert.Equal(t, errors.Is(innerErr, err1), true)

}
