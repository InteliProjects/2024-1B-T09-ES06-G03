package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashedPassword(t *testing.T) {
	password := "securepassword123"
	hashed, err := HashedPassword(password)
	if err != nil {
		t.Errorf("HashedPassword() error = %v", err)
		return
	}

	// Verificar se a senha não retorna em texto plano
	if hashed == password {
		t.Errorf("HashedPassword() should not return original password")
	}

	// Verificar se o hash pode ser comparado com a senha original corretamente
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		t.Errorf("Hashed password could not be verified with original password: %v", err)
	}
}

func TestComparePasswords(t *testing.T) {
	password := "securepassword123"
	hashed, err := HashedPassword(password)
	if err != nil {
		t.Errorf("HashedPassword() error = %v", err)
		return
	}

	// Teste de sucesso
	if !ComparePasswords(hashed, []byte(password)) {
		t.Errorf("ComparePasswords() failed to verify correct password")
	}

	// Teste de falha
	wrongPassword := "wrongpassword123"
	if ComparePasswords(hashed, []byte(wrongPassword)) {
		t.Errorf("ComparePasswords() verified wrong password as correct")
	}
}
