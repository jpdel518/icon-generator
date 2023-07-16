package service

type Service interface {
	Generate(letter string, size int) ([]byte, error)
}
