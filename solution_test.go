package solution

import (
	"fmt"
	"testing"
)

func TestGetMessage(t *testing.T) {
	_, err := fmt.Println(GetMessage())
	if err != nil {
		t.Errorf("Some error")
	}
}
