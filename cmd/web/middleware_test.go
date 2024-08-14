package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myHand myHandler
	h := NoSurf(&myHand)
	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler but is %T", v)
	}

}

func TestSessionLoad(t *testing.T) {
	var myHand myHandler
	h := SessionLoad(&myHand)
	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler but is %T", v)
	}

}

func TestWriteToConsole(t *testing.T) {
	var myHand myHandler
	h := WriteToConsole(&myHand)
	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler but is %T", v)
	}

}
