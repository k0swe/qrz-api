package qrz

import (
	"errors"
	"github.com/antihax/optional"
)

const agent = "xylo-go-1.0"

func Lookup(user *string, pw *string, call *string) (*QrzDatabase, error) {
	var config *Configuration
	config = NewConfiguration()
	client := NewAPIClient(config)

	sessResp, err := login(user, pw, client)
	if err != nil {
		return nil, err
	}

	sessionKey := sessResp.Session.Key
	if sessionKey == "" {
		return nil, errors.New(sessResp.Session.Error)
	}

	lookupResp, err := lookupInner(sessionKey, call, client)
	if err != nil {
		return nil, err
	}
	return lookupResp, nil
}

func login(user *string, pw *string, client *APIClient) (*QrzDatabase, error) {
	var req *RootGetOpts
	req = new(RootGetOpts)
	req.Username = optional.NewString(*user)
	req.Password = optional.NewString(*pw)
	req.Agent = optional.NewString(agent)
	sessResp, _, err := client.DefaultApi.RootGet(nil, req)
	if err != nil {
		return nil, err
	}
	return &sessResp, err
}

func lookupInner(sessionKey string, call *string, client *APIClient) (*QrzDatabase, error) {
	req := new(RootGetOpts)
	req.S = optional.NewString(sessionKey)
	req.Agent = optional.NewString(agent)
	req.Callsign = optional.NewString(*call)
	lookupResp, _, err := client.DefaultApi.RootGet(nil, req)
	if err != nil {
		return nil, err
	}
	return &lookupResp, nil
}
