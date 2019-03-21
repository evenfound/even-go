// Copyright (c) 2013-2017 The btcsuite developers
// Copyright (C) 2017-2019 The Even Network Developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package core

// TstAppDataDir makes the internal appDataDir function available to the test package.
func TstAppDataDir(goos, appName string, roaming bool) string {
	return appDataDir(goos, appName, roaming)
}
