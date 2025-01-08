package redditproto

import (
	"os"
	"testing"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

func TestLoad(t *testing.T) {
	expected := &UserAgent{}
	if err := prototext.Unmarshal([]byte(`
		user_agent: "test"
		client_id: "id"
		client_secret: "secret"
		username: "user"
		password: "pass"
	`), expected); err != nil {
		t.Fatalf("failed to build test expectation proto: %v", err)
	}

	testFile, err := os.CreateTemp("", "user_agent")
	if err != nil {
		t.Fatalf("failed to make test input file: %v", err)
	}

	b, err := prototext.Marshal(expected)
	if err != nil {
		t.Fatalf("failed to marshal test input: %v", err)
	}

	if _, err := testFile.Write(b); err != nil {
		t.Fatalf("failed to write test input file: %v", err)
	}

	if _, err := Load("notarealfile"); err == nil {
		t.Fatalf("wanted error returned with nonexistent file as input")
	}

	actual, err := Load(testFile.Name())
	if err != nil {
		t.Errorf("failed: %v", err)
	}

	if !proto.Equal(expected, actual) {
		t.Errorf("got %v; wanted %v", actual, expected)
	}
}
