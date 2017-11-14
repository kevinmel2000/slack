package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = strings.Index
var _ = objects.EpochTime(0)

// AuthRevokeCall is created by AuthService.Revoke method call
type AuthRevokeCall struct {
	service *AuthService
	test    bool
}

// AuthTestCall is created by AuthService.Test method call
type AuthTestCall struct {
	service *AuthService
}

// Revoke creates a AuthRevokeCall object in preparation for accessing the auth.revoke endpoint
func (s *AuthService) Revoke() *AuthRevokeCall {
	var call AuthRevokeCall
	call.service = s
	return &call
}

// Test sets the value for optional test parameter
func (c *AuthRevokeCall) Test(test bool) *AuthRevokeCall {
	c.test = test
	return c
}

// Values returns the AuthRevokeCall object as url.Values
func (c *AuthRevokeCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.test {
		v.Set("test", "true")
	}
	return v, nil
}

// Do executes the call to access auth.revoke endpoint
func (c *AuthRevokeCall) Do(ctx context.Context) error {
	const endpoint = "auth.revoke"
	v, err := c.Values()
	if err != nil {
		return err
	}
	var res struct {
		SlackResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return errors.Wrap(err, `failed to post to auth.revoke`)
	}
	if !res.OK {
		return errors.New(res.Error.String())
	}

	return nil
}

// FromValues parses the data in v and populates `c`
func (c *AuthRevokeCall) FromValues(v url.Values) error {
	var tmp AuthRevokeCall
	if raw := strings.TrimSpace(v.Get("test")); len(raw) > 0 {
		parsed, err := strconv.ParseBool(raw)
		if err != nil {
			return errors.Wrap(err, `failed to parse boolean value "test"`)
		}
		tmp.test = parsed
	}
	*c = tmp
	return nil
}

// Test creates a AuthTestCall object in preparation for accessing the auth.test endpoint
func (s *AuthService) Test() *AuthTestCall {
	var call AuthTestCall
	call.service = s
	return &call
}

// Values returns the AuthTestCall object as url.Values
func (c *AuthTestCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)
	return v, nil
}

// Do executes the call to access auth.test endpoint
func (c *AuthTestCall) Do(ctx context.Context) (*AuthTestResponse, error) {
	const endpoint = "auth.test"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*AuthTestResponse
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to auth.test`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.AuthTestResponse, nil
}

// FromValues parses the data in v and populates `c`
func (c *AuthTestCall) FromValues(v url.Values) error {
	var tmp AuthTestCall
	*c = tmp
	return nil
}
