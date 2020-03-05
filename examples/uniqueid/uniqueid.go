package uniqueid

//go:generate mockgen -source=uniqueid.go -destination=uniqueid_mock.go -package=uniqueid Generator

type Generator interface {
	Generate() string
}
