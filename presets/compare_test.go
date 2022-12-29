package presets

import (
	"fmt"
)

func compareRead(topic string, expected []byte, got []byte) error {
	if len(expected) != len(got) {
		return fmt.Errorf("%s failed, expected to read %d bytes got %d bytes", topic, len(expected), len(got))
	}
	expectedString := string(expected)
	gotString := string(got)
	if expectedString != gotString {
		return fmt.Errorf("%s failed, expected %s got %s", topic, expectedString, gotString)
	}
	return nil
}
