// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package common

// Blockchain provides a generic interface that is used to store.
// The RegisterProvider function can be used to add a new backend data storage method.
//
// This interface is divided into two distinct categories of functionality.
//
// The first category is atomic metadata storage with bucket support.
// This is accomplished through the use of database transactions.
type Blockchain interface {
	// Open the database driver type the current provider instance was created with.
	Open() error

	// Close the provider resources and free memory.
	Close()

	// Fetch balance from database provider.
	Balance(addr string) (*Balance, error)

	// Return string name of blockchain network provider.
	String() string
}

type (
	Balance struct {
		Count int
		Value float64
	}
)
