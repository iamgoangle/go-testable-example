package with_interface_test

import (
	"log"
	"testing"

	simple "github.com/iamgoangle/golfmock/with_interface"
)

type mock_doSimple struct {
	firstname, lastname string
}

// NOTES: stub the receiver method, assume that this method have http request, kafka, redis
// we don't want do with real connect.
func (m *mock_doSimple) GetFirstname() string {
	return "Mock firstname value..." + m.firstname
}

func (m *mock_doSimple) GetLastname() string {
	return m.lastname
}

func TestNewDoSimple(t *testing.T) {
	mock := &mock_doSimple{
		firstname: "Teerapong",
		lastname:  "Singthong",
	}

	got := simple.FetchDoSimple(mock)
	want := "Hi! Teerapong"

	log.Println(got)

	if got != want {
		t.Error("Not match")
	}
}
