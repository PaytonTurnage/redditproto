package redditproto

import (
	"os"

	"google.golang.org/protobuf/encoding/prototext"
)

// Load reads a user agent from a protobuffer file and returns it.
func Load(filename string) (*UserAgent, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	agent := &UserAgent{}
	return agent, prototext.Unmarshal(buf, agent)
}
