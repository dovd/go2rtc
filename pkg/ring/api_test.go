package ring

import (
	"encoding/base64"
	"encoding/json"
	"testing"
)

func TestParseAuthConfigWrappedToken(t *testing.T) {
	expected := &AuthConfig{RT: "raw-refresh-token", HID: "hardware-id"}
	b, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	config, err := parseAuthConfig(base64.StdEncoding.EncodeToString(b))
	if err != nil {
		t.Fatal(err)
	}

	if config.RT != expected.RT {
		t.Fatalf("RT = %q, want %q", config.RT, expected.RT)
	}
	if config.HID != expected.HID {
		t.Fatalf("HID = %q, want %q", config.HID, expected.HID)
	}
}

func TestParseAuthConfigRawJWTRefreshToken(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoicmVmcmVzaC10b2tlbiJ9.signature_with-url-safe-chars"

	config, err := parseAuthConfig(token)
	if err != nil {
		t.Fatal(err)
	}

	if config.RT != token {
		t.Fatalf("RT = %q, want %q", config.RT, token)
	}
	if config.HID == "" {
		t.Fatal("HID is empty")
	}
}

func TestParseAuthConfigWrappedTokenWithoutHardwareID(t *testing.T) {
	b, err := json.Marshal(&AuthConfig{RT: "raw-refresh-token"})
	if err != nil {
		t.Fatal(err)
	}

	config, err := parseAuthConfig(base64.StdEncoding.EncodeToString(b))
	if err != nil {
		t.Fatal(err)
	}

	if config.RT != "raw-refresh-token" {
		t.Fatalf("RT = %q, want raw-refresh-token", config.RT)
	}
	if config.HID == "" {
		t.Fatal("HID is empty")
	}
}
