package auth

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCreateAuthenticatorZeroAddress(t *testing.T) {
	address := common.HexToAddress("0x0000000000000000000000000000000000000000")
	// Cannot be 0x address
	_, err := CreateAuthenticator(address)
	require.Equal(t, err != nil, true)
}

// Test signing with authenticator
func TestAuthenticatorSign(t *testing.T) {
	address, _, privateKey, _ := CreateWallet()
	auth, _ := CreateAuthenticator(address)
	_, err := auth.Sign([]byte("hello world"), privateKey)
	require.Equal(t, err, nil)
}

// Test verifying a signature
func TestAuthenticatorVerifyPass(t *testing.T) {
	address, _, privateKey, _ := CreateWallet()
	auth, _ := CreateAuthenticator(address)
	msg := []byte("hello world")
	sigHex, _ := auth.Sign(msg, privateKey)
	require.Equal(t, auth.Verify(sigHex, msg), true)
}

