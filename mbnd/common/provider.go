// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package common

import (
	"fmt"

	"github.com/go-errors/errors"
)

// Provider defines a structure for backend providers.
type Provider struct {
	// Token is the identifier used to uniquely identify a specific blockchain.
	Token string

	// Init is the function that will be invoked with all user-specified arguments to open the connections.
	Init func(network string) (Blockchain, error)
}

// driverList holds all of the registered blockchain nets.
var providers = make(map[string]*Provider)

// RegisterProvider adds a backend database driver to available interfaces.
// ErrDbTypeRegistered will be returned if the database type for the driver has
// already been registered.
func RegisterProvider(provider Provider) error {

	if _, exists := providers[provider.Token]; exists {
		return errors.New(fmt.Sprintf("provider %q is already registered", provider.Token))
	}

	providers[provider.Token] = &provider

	return nil
}

// SupportedProvider returns boolean if specified name that have been registered and supported.
func SupportedProvider(name string) bool {

	for _, drv := range providers {
		if drv.Token == name {
			return true
		}
	}

	return false
}

// SupportingProviders returns a slice of strings that represent  providers that have been registered and are supported.
func SupportingProviders() []string {

	supported := make([]string, 0, len(providers))

	for _, drv := range providers {
		supported = append(supported, drv.Token)
	}

	return supported
}

// Init connect to an existing blockchain for the specified type.
func Init(token string, network string) (Blockchain, error) {

	provider, exists := providers[token]

	if !exists {
		return nil, errors.New(fmt.Sprintf("provider %q is not registered", token))
	}

	return provider.Init(network)
}
