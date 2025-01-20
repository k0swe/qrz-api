package qrz

import (
	"context"
	"errors"
)

const agent = "k0swe-go-1.0"

var cachedUser = ""
var cachedSession = ""

func Lookup(ctx context.Context, user *string, pw *string, call *string) (*QRZDatabase, error) {
	config := NewConfiguration()
	config.UserAgent = agent
	client := NewAPIClient(config)

	sessionKey, err := login(ctx, user, pw, client)
	if err != nil {
		return nil, err
	}

	lookupResp, err := lookupInner(ctx, sessionKey, call, client)
	if err != nil {
		cachedSession = ""
		return nil, err
	}
	return lookupResp, nil
}

func login(ctx context.Context, user *string, pw *string, client *APIClient) (string, error) {
	if cachedUser == *user && cachedSession != "" {
		return cachedSession, nil
	}
	req := client.DefaultAPI.RootGet(ctx).Username(*user).Password(*pw).Agent(agent)
	sessResp, _, err := req.Execute()
	if err != nil {
		return "", err
	}
	sessionKey := *sessResp.Session.Key
	if sessionKey == "" {
		return "", errors.New(*sessResp.Session.Error)
	}
	cachedUser = *user
	cachedSession = sessionKey
	return sessionKey, err
}

func lookupInner(ctx context.Context, sessionKey string, call *string, client *APIClient) (*QRZDatabase, error) {
	req := client.DefaultAPI.RootGet(ctx).S(sessionKey).Agent(agent)
	req = req.Callsign(*call)
	lookupResp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return lookupResp, nil
}
