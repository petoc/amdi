package amdi

import (
	"testing"
)

func TestIs(t *testing.T) {
	cs := map[string]bool{
		"Watch1,1":  true,
		"iPad1,1":   true,
		"iPhone1,1": true,
		"iPod1,1":   true,
		"iPod1,2":   false,
	}
	for c, e := range cs {
		g := Is(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}

func TestName(t *testing.T) {
	cs := map[string]string{
		"Watch1,1":  "Watch 38mm case",
		"iPad1,1":   "iPad",
		"iPhone1,1": "iPhone",
		"iPod1,1":   "iPod",
		"iPod1,2":   "",
	}
	for c, e := range cs {
		g := Name(c)
		if g != e {
			t.Errorf("%s: expected %v, got %v", c, e, g)
			return
		}
	}
}
