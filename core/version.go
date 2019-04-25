// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package core

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// semanticAlphabet
	semanticAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	// These constants define the application semantic and follow the semantic versioning 2.0.0 spec (http://semver.org/).
	appMajor uint = 0
	appMinor uint = 15
	appPatch uint = 0

	// appPreRelease MUST only contain characters from semanticAlphabet per the semantic versioning spec.
	appPreRelease = "alfa"
)

// It MUST only contain characters from semanticAlphabet per the semantic versioning spec.
var appBuild string

// Print the core version per the semantic versioning 2.0.0 spec (http://semver.org/).
func SprintVersion(tmpl ...interface{}) string {
	var format string
	if tmpl == nil {
		format = "Even Network core version %s"
	} else {
		format = strings.Trim(fmt.Sprint(tmpl), "[]")
	}
	return fmt.Sprintf(format, version())
}

// semantic returns the application semantic as a properly formed string
// per the semantic versioning 2.0.0 spec (http://semver.org/).
func version() string {

	// Start with the major, minor, and patch versions.
	version := fmt.Sprintf("%d.%d.%d", appMajor, appMinor, appPatch)

	// Append pre-release s if there is one.
	// The hyphen called for by the s versioning spec is automatically appended and should
	// not be contained in the pre-release string.
	// The pre-release s is not appended if it contains invalid characters.
	preRelease := normalizeVerString(appPreRelease)
	if preRelease != "" {
		version = fmt.Sprintf("%s-%s", version, preRelease)
	}

	// Append build metadata if there is any.
	// The plus called for by the s versioning spec is automatically appended and should
	// not be contained in the build metadata string.
	// The build metadata string is not appended if it contains invalid characters.
	build := normalizeVerString(appBuild)
	if build != "" {
		version = fmt.Sprintf("%s+%s", version, build)
	}

	return version
}

// normalizeVerString returns the passed string stripped of all characters which are not valid according
// to the semantic versioning guidelines for pre-release semantic and build metadata strings.
// In particular they MUST only contain characters in semanticAlphabet.
func normalizeVerString(str string) string {

	var result bytes.Buffer

	for _, r := range str {
		if strings.ContainsRune(semanticAlphabet, r) {
			// Ignoring the error here since it can only fail if the the system is out of memory and there are much
			// bigger issues at that point.
			_, _ = result.WriteRune(r)
		}
	}

	return result.String()
}
