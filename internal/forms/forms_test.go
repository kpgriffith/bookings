package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	has := form.Has("a")
	if has {
		t.Error("says it has the field when it doesn't")
	}

	postedData = url.Values{}
	postedData.Add("a", "abc")
	form = New(postedData)
	has = form.Has("a")
	if !has {
		t.Error("says it doesn't have the field when it does")
	}

}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	isMinLength := form.MinLength("a", 3)
	if isMinLength {
		t.Error("found non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("a_field", "a")
	form = New(postedData)
	isMinLength = form.MinLength("a_field", 3)
	if isMinLength {
		t.Error("says it's min length when it isn't")
	}

	postedData = url.Values{}
	postedData.Add("b_field", "abc")
	form = New(postedData)
	isMinLength = form.MinLength("b_field", 3)
	if !isMinLength {
		t.Error("says it's not min length when it is")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "a")
	form := New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("says it's valid when it shouldn't be")
	}

	postedData = url.Values{}
	postedData.Add("a", "a@a.com")
	form = New(postedData)
	form.IsEmail("a")
	if !form.Valid() {
		t.Error("says its invalid when it shouldn't be")
	}
}
