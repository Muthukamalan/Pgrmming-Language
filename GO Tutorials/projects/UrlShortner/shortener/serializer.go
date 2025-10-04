package shortener

type RedirectSerializer interface {
	Decode(ip []byte) (*Redirect, error)
	Encode(ip *Redirect) ([]byte, error)
}
