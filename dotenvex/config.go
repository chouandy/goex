package dotenvex

var filePrefix = ".env"

var encryptedFileExt = ".enc"

// SetFilePrefix set file prefix
func SetFilePrefix(s string) {
	filePrefix = s
}

// SetEncryptedFileExt set encrypted file ext
func SetEncryptedFileExt(s string) {
	encryptedFileExt = s
}
