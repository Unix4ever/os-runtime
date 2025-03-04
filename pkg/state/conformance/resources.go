// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package conformance

import (
	"fmt"

	"github.com/talos-systems/os-runtime/pkg/resource"
)

// PathResourceType is the type of PathResource.
const PathResourceType = resource.Type("os/path")

// PathResource represents a path in the filesystem.
//
// Resource ID is the path.
type PathResource struct {
	md resource.Metadata
}

// NewPathResource creates new PathResource.
func NewPathResource(ns resource.Namespace, path string) *PathResource {
	r := &PathResource{
		md: resource.NewMetadata(ns, PathResourceType, path, resource.VersionUndefined),
	}
	r.md.BumpVersion()

	return r
}

// Metadata implements resource.Resource.
func (path *PathResource) Metadata() *resource.Metadata {
	return &path.md
}

// Spec implements resource.Resource.
func (path *PathResource) Spec() interface{} {
	return nil
}

func (path *PathResource) String() string {
	return fmt.Sprintf("PathResource(%q)", path.md.ID())
}

// DeepCopy implements resource.Resource.
func (path *PathResource) DeepCopy() resource.Resource {
	return &PathResource{
		md: path.md,
	}
}
