// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package etcd

// StateOptions configure inmem.State.
type StateOptions struct{}

// StateOption applies settings to StateOptions.
type StateOption func(options *StateOptions)

// DefaultStateOptions returns default value of StateOptions.
func DefaultStateOptions() StateOptions {
	return StateOptions{}
}
