package shortener

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo IRedirectRepository
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errors.Join(ErrRedirectNotFound, err)
	} else {
		redirect.Code = shortid.MustGenerate()
		redirect.CreatedAt = time.Now().UTC().Unix()
		return r.redirectRepo.Store(redirect)
	}
}
