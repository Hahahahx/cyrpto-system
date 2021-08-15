package action

type GetOptions struct {
	CID     string
	Newname string
	Decrypt bool
}

type AddOptions struct {
	Filename string
	Encrypt  bool
}

type DecryptOptions struct {
	Filename string
	Key      []byte
	Newname  string
}

type EncryptOptions struct {
	Filename string
}
