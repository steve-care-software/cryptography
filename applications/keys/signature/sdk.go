package signature

import (
    "github.com/steve-care-software/cryptography/domain/hash"
)

// Application represents the signature application
type Application interface {
    New() (*hash.Hash, error)
    Derive(pkContext hash.Hash) ([]byte, error)
    Sign(msg []byte, pkContext hash.Hash) ([]byte, error)
    Validate(sig []byte, msg []byte) bool
    RingSign(msg []byte, pkContext hash.Hash, ring []hash.Hash) ([]byte, error)
    RingValidate(sig []byte, msg []byte) bool
}
