/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodes

import (
	v1alpha1nodes "github.com/pokt-network/pocket-operator/apis/nodes/v1alpha1"
	v1alpha1pocketset "github.com/pokt-network/pocket-operator/apis/nodes/v1alpha1/pocketset"
)

// Code generated by operator-builder. DO NOT EDIT.

// PocketSetLatestGroupVersion returns the latest group version object associated with this
// particular kind.
var PocketSetLatestGroupVersion = v1alpha1nodes.GroupVersion

// PocketSetLatestSample returns the latest sample manifest associated with this
// particular kind.
var PocketSetLatestSample = v1alpha1pocketset.Sample(false)
