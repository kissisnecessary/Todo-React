package workshop

type SM2 interface {
	Encrypt(msg []byte) ([]byte, error)
	Decrypt(encrypted []by