package modelv1

type Controller interface {
	Load() error
	Start()
	Stop()

	SaveUser() (*User, error)
}
