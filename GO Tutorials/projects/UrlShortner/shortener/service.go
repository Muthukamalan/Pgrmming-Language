package shortener

type IRedirectService interface {
	Find(code string) (*Redirect, error)
	Store(r *Redirect) error
}
