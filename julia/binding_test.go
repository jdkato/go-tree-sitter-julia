package julia

import (
	"testing"
)

func TestGrammar(t *testing.T) {
	lang := GetLanguage()
	if lang == nil {
		t.Fatal("Failed to get language")
	}
}
