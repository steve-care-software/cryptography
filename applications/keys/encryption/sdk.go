package encryption

import (
    "github.com/steve-care-software/cryptography/domain/hash"
)

// Application represents the encryption application
type Application interface {
    New(bitrate uint) (*hash.Hash, error)
    AddFromBytes(bytes []byte) (*hash.Hash, error)
    AddFromBase64(base string) (*hash.Hash, error)
    Derive(pkContext hash.Hash) ([]byte, error)
    Encrypt(msg []byte, pkContext hash.Hash) ([]byte, error)
    Decrypt(cipher []byte, pubKeyContext hash.Hash) ([]byte, error)
    FetchToBytes(context hash.Hash) ([]byte, error)
    FetchToBase64(context hash.Hash) (string, error)
}
