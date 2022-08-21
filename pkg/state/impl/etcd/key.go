// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package etcd

import (
	"fmt"
	"net/url"

	"github.com/cosi-project/runtime/pkg/resource"
)

const (
	etcdKeyPrefix = "cosi"
)

// etcdKeyFromPointer builds a key for the resource with the given pointer.
func etcdKeyFromPointer(pointer resource.Pointer) string {
	prefix := etcdKeyPrefixFromKind(pointer)
	idEscaped := url.QueryEscape(pointer.ID())

	return fmt.Sprintf("%s%s", prefix, idEscaped)
}

// etcdKeyPrefixFromKind returns a key prefix for the given resource kind.
func etcdKeyPrefixFromKind(kind resource.Kind) string {
	nsEscaped := url.QueryEscape(kind.Namespace())
	typeEscaped := url.QueryEscape(kind.Type())

	return fmt.Sprintf("/%s/%s/%s/", etcdKeyPrefix, nsEscaped, typeEscaped)
}
