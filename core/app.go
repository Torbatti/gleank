package core

type App interface {
	DataDir() string
	IsDev() bool
}
