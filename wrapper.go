package qrz

import (
	"context"
	"errors"
	"github.com/antihax/optional"
)

const agent = "xylo-go-1.0"

func Lookup(user *string, pw *string, call *string) (*QrzDatabase, error) {
	config := NewConfiguration()
	config.UserAgent = agent
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
	req := new(RootGetOpts)
	req.Username = optional.NewString(*user)
	req.Password = optional.NewString(*pw)
	req.Agent = optional.NewString(agent)
	sessResp, _, err := client.DefaultApi.RootGet(context.TODO(), req)
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
	lookupResp, _, err := client.DefaultApi.RootGet(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	return &lookupResp, nil
}
