package shortener

type IRedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(r *Redirect) error
}
