package scan

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("k8s")
	if client.Scanner != "k8s" {
		t.Errorf("expected k8s, got %s", client.Scanner)
	}
}