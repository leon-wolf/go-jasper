package gojasper

import "errors"

func New(url, username, password string, pro bool) (*API, error) {
	if len(url) != 0 && len(username) != 0 && len(password) != 0 {
		a := API{
			URL:      url,
			Username: username,
			Password: password,
			Pro:      pro,
		}
		_, err := a.ServerInfo()
		if err != nil {
			return nil, err
		}
		return &a, nil
	}
	return nil, errors.New("empty parameter")
}
