package redditproto

import (
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestLoad(t *testing.T) {
	expected := &UserAgent{}
	if err := proto.UnmarshalText(`
		user_agent: "test"
		client_id: "id"
		client_secret: "secret"
		username: "user"
		password: "pass"
	`, expected); err != nil {
		t.Errorf("failed to build test expectation proto: %v", err)
	}

	testFile, err := ioutil.TempFile("", "user_agent")
	if err != nil {
		t.Errorf("failed to make test input file: %v", err)
	}

	if err := proto.MarshalText(testFile, expected); err != nil {
		t.Errorf("failed to write test input file: %v", err)
	}

	if _, err := Load("notarealfile"); err == nil {
		t.Errorf("wanted error returned with nonexistent file as input")
	}

	actual, err := Load(testFile.Name())
	if err != nil {
		t.Errorf("failed: %v", err)
	}

	if !proto.Equal(expected, actual) {
		t.Errorf("got %v; wanted %v", actual, expected)
	}
}
