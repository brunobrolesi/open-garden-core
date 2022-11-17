package gateway

type Hash = string

type HashService interface {
	GenerateHash(s string) (Hash, error)
	CompareStringAndHash(s, hash string) error
}
