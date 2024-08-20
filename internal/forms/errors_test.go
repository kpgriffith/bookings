package forms

import (
	"net/url"
	"testing"
)

func TestErrors_Get(t *testing.T) {

	postedData := url.Values{}
	form := New(postedData)

	errString := form.Errors.Get("field1")
	if len(errString) != 0 {
		t.Error("found an error when it shouldn't have")
	}

	form.Errors.Add("field2", "error on field 2")
	errString = form.Errors.Get("field2")
	if len(errString) == 0 {
		t.Error("did't find an error it should have")
	}

}
