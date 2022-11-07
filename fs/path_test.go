package fs

import (
	"testing"
)

func TestPathToFormattedName(t *testing.T) {
	// shallow
	name := PathToFormattedName(".env.hello")

	if name != "hello" {
		t.Error("expected 'hello' got", name)
	}

	// nested
	name = PathToFormattedName("local/.env.hello")

	if name != "local/hello" {
		t.Error("expected 'local/hello' got", name)
	}
}

func TestPathFromFormattedName(t *testing.T) {
	// shallow
	path := PathFromFormattedName("hello")

	if path != ".env.hello" {
		t.Error("expected '.env.hello' got", path)
	}

	// nested
	path = PathFromFormattedName("local/hello")

	if path != "local/.env.hello" {
		t.Error("expected 'local/.env.hello' got", path)
	}
}
