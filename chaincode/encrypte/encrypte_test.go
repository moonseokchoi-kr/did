package encrypte

import (
	"strconv"
	"testing"
)

func TestDIDid(t *testing.T) {
	want := "Hello, world."
	if got := "Hello, world."; got != want {
		t.Errorf("getSpecificID = %q, want %q", got, want)
	}
}

func TestGenerateKey(t *testing.T) {
	private, public, sig := MakeECDSAKey("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	t.Logf("privatekey : %q, \n publickey : %q, \n siginature : %q", private, public, sig)
}

func TestVerify(t *testing.T) {
	pubkey := "13n4s5tFAmoCYHLsnJ9k1nspszbuQgvjaFrmJ8cSbfLmHDGNDkc69XCExX9PpbDBLA25VK2GsvYXvXEi9xr1DWEbVfUJu8u"
	signature := "124GJCUhBd4Nofx3bbN6aHuyjMcTeaBmtPtGar1j9iXYw1ZVzGMwXc5CGQP4nz1RBvk64sS7ecJCKdZR6YgbtSHzGcbFWfAtFqRAwjon"
	msg := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	got := Verify(pubkey, signature, msg)
	want := true
	if got != want {
		t.Errorf("verification failed! got: %q", strconv.FormatBool(got))
	}
}

func TestGetJWT(t *testing.T) {
	got, err := GetJWT()
	if err != nil {
		t.Error(err)
	}
	want := "no jwt"
	if got != want {
		t.Errorf("jwt string: %q", got)
	}
}

func TestDecodeJWT(t *testing.T) {

}
