// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package resource

import (
	"fmt"
	"strconv"

	"github.com/siderolabs/go-pointer"
)

// Version of a resource.
type Version struct {
	// make versions uncomparable with equality operator
	_ [0]func()

	*uint64
}

// Special version constants.
var (
	VersionUndefined = Version{}
)

const undefinedVersion = "undefined"

// Value returns the underlying version number. Returns 0 if the version is undefined.
func (v Version) Value() uint64 {
	return pointer.SafeDeref(v.uint64)
}

// Next returns a new incremented version.
func (v Version) Next() Version {
	return Version{
		uint64: pointer.To(pointer.SafeDeref(v.uint64) + 1),
	}
}

func (v Version) String() string {
	if v.uint64 == nil {
		return undefinedVersion
	}

	return strconv.FormatUint(*v.uint64, 10)
}

// Equal compares versions.
func (v Version) Equal(other Version) bool {
	if v.uint64 == nil || other.uint64 == nil {
		return v.uint64 == nil && other.uint64 == nil
	}

	return *v.uint64 == *other.uint64
}

// ParseVersion from string representation.
func ParseVersion(ver string) (Version, error) {
	if ver == undefinedVersion {
		return VersionUndefined, nil
	}

	intVersion, err := strconv.ParseInt(ver, 10, 64)
	if err != nil {
		return VersionUndefined, fmt.Errorf("error parsing version: %w", err)
	}

	return Version{
		uint64: pointer.To(uint64(intVersion)),
	}, nil
}

// NewVersion returns a new version for the given version number.
func NewVersion(ver uint64) Version {
	return Version{
		uint64: pointer.To(ver),
	}
}
