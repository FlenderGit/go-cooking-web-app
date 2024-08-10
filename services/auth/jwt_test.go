package auth

import (
	"testing"
)

type ClaimsTest struct {
	ID    int
	Login string
}

var listClaimsValid = []ClaimsTest{
	{1, "admin"},
	{2, "user"},
	{3, "test"},
}

func TestGenerateToken(t *testing.T) {

	for _, claimsTest := range listClaimsValid {
		token, err := GenerateToken(claimsTest.ID, claimsTest.Login)

		if err != nil {
			t.Errorf("Error while generating token: %v", err)
		}

		if token == "" {
			t.Errorf("Token is empty")
		}
	}

}

func TestParseToken(t *testing.T) {

	for _, claimsTest := range listClaimsValid {
		token, _ := GenerateToken(claimsTest.ID, claimsTest.Login)
		claims, err := ParseToken(token)

		if err != nil {
			t.Errorf("Error while parsing token: %v", err)
		}

		if claims.ID != claimsTest.ID {
			t.Errorf("ID is not the same")
		}

		if claims.Login != claimsTest.Login {
			t.Errorf("Username is not the same")
		}
	}
}

func TestInvalidToken(t *testing.T) {
	token := "invalid_token"
	_, err := ParseToken(token)

	if err == nil {
		t.Errorf("Token should be invalid")
	}
}

func TestGenerateTokenInvalidIdNegative(t *testing.T) {
	_, err := GenerateToken(-1, "admin")

	if err == nil {
		t.Errorf("Token should be invalid")
	}
}

func TestGenerateTokenLoginEmpty(t *testing.T) {
	_, err := GenerateToken(1, "")

	if err == nil {
		t.Errorf("Token should be invalid")
	}
}
