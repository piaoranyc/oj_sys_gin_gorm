package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(T *testing.T) {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.NewV4().String()
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling

}
