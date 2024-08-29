package db

import "testing"

func Test_Connect(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Error(err)
	}
}
