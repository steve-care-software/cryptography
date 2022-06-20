package password

// Application represents the password application
type Application interface {
    Encrypt(msg []byte, password []byte) ([]byte, error)
    Decrypt(cipher []byte, password []byte) ([]byte, error)
}
