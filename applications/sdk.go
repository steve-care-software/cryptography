package applications

import (
    "github.com/steve-care-software/cryptography/applications/hash"
    "github.com/steve-care-software/cryptography/applications/keys"
    "github.com/steve-care-software/cryptography/applications/password"
)

// Application represents the application
type Application interface {
    Hash() hash.Application
    Keys() keys.Application
    Password() password.Application
}
