package hash

import (
    "github.com/steve-care-software/cryptography/domain/hash"
)

// Application represents an hash application
type Application interface {
    FromBytes(data []byte) (*hash.Hash, error)
    FromMultiBytes(data [][]byte) (*hash.Hash, error)
    FromBase64(base string) (*hash.Hash, error)
}
