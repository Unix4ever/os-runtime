// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/talos-systems/os-runtime/pkg/resource"
	"github.com/talos-systems/os-runtime/pkg/state"
	"github.com/talos-systems/os-runtime/pkg/state/impl/inmem"
)

func TestErrors(t *testing.T) {
	t.Parallel()

	assert.True(t, state.IsNotFoundError(inmem.ErrNotFound(resource.NewMetadata("ns", "a", "b", resource.VersionUndefined))))
	assert.True(t, state.IsConflictError(inmem.ErrAlreadyExists(resource.NewMetadata("ns", "a", "b", resource.VersionUndefined))))
	assert.True(t, state.IsConflictError(inmem.ErrVersionConflict(resource.NewMetadata("ns", "a", "b", resource.VersionUndefined), resource.VersionUndefined, resource.VersionUndefined)))
	assert.True(t, state.IsConflictError(inmem.ErrPendingFinalizers(resource.NewMetadata("ns", "a", "b", resource.VersionUndefined))))
}
