package pivotal_test

import (
	"testing"

	"github.com/danimal141/pivotal-go/pivotal"
)

func TestClient(t *testing.T) {
	const token = "some-token"
	cli, err := pivotal.NewClient(token)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if cli.Token != token {
		t.Errorf("Expected %v, got %v", token, cli.Token)
	}
}
