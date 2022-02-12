package solution

import (
	"testing"
  
	"github.com/kyokomi/emoji/v2"
  )
  
  func TestHello(t *testing.T) {
	got := GetMessage()
	want := emoji.Sprint("Hello 🗺️ !")
  
	if got != want {
	  t.Errorf("got %q want %q", got, want)
	}
  }