package oidctestserver

import (
	"github.com/jonas-plum/maut/oidctestserver/internal"
)

func NativeClient(id string, redirectURIs ...string) *internal.Client {
	return internal.NativeClient(id, redirectURIs...)
}

func WebClient(id, secret string, redirectURIs ...string) *internal.Client {
	return internal.WebClient(id, secret, redirectURIs...)
}
