package auth

import (
	"regexp"
	"errors"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	accounts "github.com/ethereum/go-ethereum/accounts"
	hexutil "github.com/ethereum/go-ethereum/common/hexutil"
	crypto "github.com/ethereum/go-ethereum/crypto"
)

// CreateWallet - Generate new public, private key pair
// This function is mostly used for testing
func CreateWallet() (common.Address, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}, "", "", err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, "", "", errors.New("CreateWallet: error in public key creation")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyString := hexutil.Encode(publicKeyBytes)[4:]

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	
	return address, publicKeyString, privateKeyString, nil
}

// Message - A struct for sending a signed message
type Message struct {
	Method string `json:"method"`
	RequestPath string `json:"requestPath"`
	Body string `json:"body"`
	Timestamp int64 `json:"timestamp"`
}

// CreateMessage - Create `Message` instance
// Arguments:
// 	method - GET or POST
// 	requestPath - URI in the path
// 	body - Stringified of the body
// 	timestamp - UNIX timestamp
func CreateMessage(
	method string,
	requestPath string,
	body string,
	timestamp int64,
) (*Message, error) {
	message := &Message{
		Method: method,
		RequestPath: requestPath,
		Body: body,
		Timestamp: timestamp,
	}
	return message, nil
}

// Authenticator - A struct for checking user signatures
type Authenticator struct {
	address common.Address
}

// CreateAuthenticator - Create `Authenticator` instance
// Arguments:
// 	address - Address for the underlying asset
func CreateAuthenticator(address common.Address) (*Authenticator, error) {
	auth := &Authenticator{address: address}
	if auth.ValidAddress() {
		return auth, nil
	}
	return nil, errors.New("CreateAuthenticator: invalid address")
}

// ValidAddress - Checks if an address is valid
func (auth *Authenticator) ValidAddress() bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(auth.address.Hex()) {
		return false
	}
	if auth.address.Hex() == "0x0000000000000000000000000000000000000000" {
		return false
	}
	return true
}

// Verify - Function to verify a signature
// Arguments:
// 	sigHex - Signature to be verified
// 	msg - The message that should be sent
// Notes:
// 	Taken from https://gist.github.com/dcb9/385631846097e1f59e3cba3b1d42f3ed#file-eth_sign_verify-go
func (auth *Authenticator) Verify(sigHex string, msg []byte) bool {
	// Parse the hexcode for byte string of signature
	sig := hexutil.MustDecode(sigHex)

	// Compute keccak hash of message
	msg = accounts.TextHash(msg)

	// Transform yellow paper V from 27/28 to 0/1
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 
	}

	// Derive public key from signature
	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	// Derive address from public key
	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	// True if the recovered address is the expected one
	return auth.address.Hex() == recoveredAddr.Hex()
}

// Sign - Function to sign with private key
// Arguments:
// 	msg - The message that should be sent
// 	privateKey - Private key for the user. Must be related to the public key 
func (auth *Authenticator) Sign(msg []byte, privateKey string) (string, error) {
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	msg = accounts.TextHash(msg)
	sig, err := crypto.Sign(msg, privateKeyEcdsa)
	if err != nil {
		return "", err
	}
	sigHex := hexutil.Encode(sig)
	return sigHex, nil
}
