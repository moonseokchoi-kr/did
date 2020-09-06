package chaincode

import (
	"testing"
)

func TestDIDid(t *testing.T) {
	want := "Hello, world."
	if got := "Hello, world."; got != want {
		t.Errorf("getSpecificID = %q, want %q", got, want)
	}
}

func TestGenerateKey(t *testing.T) {
	private, public, sig := makeECDSAKey()
	t.Errorf("privatekey : %q, \n publickey : %q, \n siginature : %q", private, public, sig)
}
