// Copyright PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/federation/pingcap/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeBackups implements VolumeBackupInterface
type FakeVolumeBackups struct {
	Fake *FakeFederationV1alpha1
	ns   string
}

var volumebackupsResource = v1alpha1.SchemeGroupVersion.WithResource("volumebackups")

var volumebackupsKind = v1alpha1.SchemeGroupVersion.WithKind("VolumeBackup")

// Get takes name of the volumeBackup, and returns the corresponding volumeBackup object, and an error if there is any.
func (c *FakeVolumeBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumebackupsResource, c.ns, name), &v1alpha1.VolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeBackup), err
}

// List takes label and field selectors, and returns the list of VolumeBackups that match those selectors.
func (c *FakeVolumeBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumebackupsResource, volumebackupsKind, c.ns, opts), &v1alpha1.VolumeBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VolumeBackupList{ListMeta: obj.(*v1alpha1.VolumeBackupList).ListMeta}
	for _, item := range obj.(*v1alpha1.VolumeBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeBackups.
func (c *FakeVolumeBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumebackupsResource, c.ns, opts))

}

// Create takes the representation of a volumeBackup and creates it.  Returns the server's representation of the volumeBackup, and an error, if there is any.
func (c *FakeVolumeBackups) Create(ctx context.Context, volumeBackup *v1alpha1.VolumeBackup, opts v1.CreateOptions) (result *v1alpha1.VolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumebackupsResource, c.ns, volumeBackup), &v1alpha1.VolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeBackup), err
}

// Update takes the representation of a volumeBackup and updates it. Returns the server's representation of the volumeBackup, and an error, if there is any.
func (c *FakeVolumeBackups) Update(ctx context.Context, volumeBackup *v1alpha1.VolumeBackup, opts v1.UpdateOptions) (result *v1alpha1.VolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumebackupsResource, c.ns, volumeBackup), &v1alpha1.VolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeBackup), err
}

// Delete takes name of the volumeBackup and deletes it. Returns an error if one occurs.
func (c *FakeVolumeBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(volumebackupsResource, c.ns, name, opts), &v1alpha1.VolumeBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumebackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VolumeBackupList{})
	return err
}

// Patch applies the patch and returns the patched volumeBackup.
func (c *FakeVolumeBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumebackupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeBackup), err
}
