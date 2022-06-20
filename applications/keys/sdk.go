package keys

import (
    "github.com/steve-care-software/cryptography/applications/keys/encryption"
    "github.com/steve-care-software/cryptography/applications/keys/signature"
)

// Application represents the keys application
type Application interface {
    Encryption() encryption.Application
    Signature() signature.Application
}
