package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/random", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Errorf("expected form to be valid, got %v", isValid)
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/random", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("expected form to be invalid since required fields are missing, got valid")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "a")
	postData.Add("c", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("not_exist")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form shows does not have field when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.MinLength("a", 3)
	if form.Valid() {
		t.Error("form shows has field when it does not")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postData = url.Values{}
	postData.Add("a", "ab")

	form = New(postData)

	form.MinLength("a", 3)
	if form.Valid() {
		t.Error("form shows has field when it does not")
	}

	postData = url.Values{}
	postData.Add("a", "abc")
	form = New(postData)
	form.MinLength("a", 1)

	if !form.Valid() {
		t.Error("shows min length of 1 is not met when it is")
	}
	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error, but get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")

	if form.Valid() {
		t.Error("form shows valid email for not existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address when we should not have")
	}
}
