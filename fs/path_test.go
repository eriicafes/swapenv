package fs

import (
	"testing"
)

func TestFormatPath(t *testing.T) {
	// shallow
	name := FormatPath(".env.hello")

	if name != "hello" {
		t.Error("expected 'hello' got", name)
	}

	// nested
	name = FormatPath("local/.env.hello")

	if name != "local/hello" {
		t.Error("expected 'local/hello' got", name)
	}
}

func TestNormalizePath(t *testing.T) {
	// shallow
	path := NormalizePath("hello")

	if path != ".env.hello" {
		t.Error("expected '.env.hello' got", path)
	}

	// nested
	path = NormalizePath("local/hello")

	if path != "local/.env.hello" {
		t.Error("expected 'local/.env.hello' got", path)
	}
}
