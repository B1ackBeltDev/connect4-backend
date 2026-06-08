package sessions

type SessionsStorage interface {
	Get(key string) (string, error)
	Store(key, value string) error
}
