package gflights

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/browserutils/kooky"
	"github.com/tobyrushton/gflights/internal/syncmap"
)

//go:generate go tool counterfeiter -generate

//counterfeiter:generate -o internal/fakes . HTTPDoer
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type SessionOptions func(s *Session)

type Session struct {
	Cities syncmap.Map[string, string]

	client  HTTPDoer
	cookies []string
}

func getCookies(res *http.Response) ([]string, error) {
	var cookies []string
	if setCookie, ok := res.Header["Set-Cookie"]; ok {
		for _, c := range setCookie {
			cookies = append(cookies, strings.Split(c, ";")[0])
		}
		return cookies, nil
	}
	return nil, fmt.Errorf("could not find the 'Set-Cookie' header in the initialization response")
}

func New(opts ...SessionOptions) (*Session, error) {
	client := &http.Client{}

	res, err := client.Get("https://www.google.com")
	if err != nil {
		return nil, fmt.Errorf("new session: err sending request to www.google.com: %v", err)
	}

	cookies, err := getCookies(res)
	if err != nil {
		return nil, fmt.Errorf("new session: err getting cookies: %v", err)
	}

	ctx := context.Background()
	GOOGLE_ABUSE_EXEMPTION, err := kooky.ReadCookies(ctx, kooky.Valid, kooky.DomainHasSuffix(`google.com`), kooky.Name(`GOOGLE_ABUSE_EXEMPTION`))
	if err != nil {
		return nil, fmt.Errorf("new session: err reading GOOGLE_ABUSE_EXEMPTION cookie: %v", err)
	}

	if len(GOOGLE_ABUSE_EXEMPTION) == 1 {
		cookies = append(cookies, GOOGLE_ABUSE_EXEMPTION[0].Value)
	}

	return &Session{
		client:  client,
		cookies: cookies,
	}, nil
}

func WithClient(client HTTPDoer) SessionOptions {
	return func(s *Session) {
		s.client = client
	}
}
