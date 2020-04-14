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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TidbClusterAutoScalerLister helps list TidbClusterAutoScalers.
type TidbClusterAutoScalerLister interface {
	// List lists all TidbClusterAutoScalers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TidbClusterAutoScaler, err error)
	// TidbClusterAutoScalers returns an object that can list and get TidbClusterAutoScalers.
	TidbClusterAutoScalers(namespace string) TidbClusterAutoScalerNamespaceLister
	TidbClusterAutoScalerListerExpansion
}

// tidbClusterAutoScalerLister implements the TidbClusterAutoScalerLister interface.
type tidbClusterAutoScalerLister struct {
	indexer cache.Indexer
}

// NewTidbClusterAutoScalerLister returns a new TidbClusterAutoScalerLister.
func NewTidbClusterAutoScalerLister(indexer cache.Indexer) TidbClusterAutoScalerLister {
	return &tidbClusterAutoScalerLister{indexer: indexer}
}

// List lists all TidbClusterAutoScalers in the indexer.
func (s *tidbClusterAutoScalerLister) List(selector labels.Selector) (ret []*v1alpha1.TidbClusterAutoScaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TidbClusterAutoScaler))
	})
	return ret, err
}

// TidbClusterAutoScalers returns an object that can list and get TidbClusterAutoScalers.
func (s *tidbClusterAutoScalerLister) TidbClusterAutoScalers(namespace string) TidbClusterAutoScalerNamespaceLister {
	return tidbClusterAutoScalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TidbClusterAutoScalerNamespaceLister helps list and get TidbClusterAutoScalers.
type TidbClusterAutoScalerNamespaceLister interface {
	// List lists all TidbClusterAutoScalers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TidbClusterAutoScaler, err error)
	// Get retrieves the TidbClusterAutoScaler from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TidbClusterAutoScaler, error)
	TidbClusterAutoScalerNamespaceListerExpansion
}

// tidbClusterAutoScalerNamespaceLister implements the TidbClusterAutoScalerNamespaceLister
// interface.
type tidbClusterAutoScalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TidbClusterAutoScalers in the indexer for a given namespace.
func (s tidbClusterAutoScalerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TidbClusterAutoScaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TidbClusterAutoScaler))
	})
	return ret, err
}

// Get retrieves the TidbClusterAutoScaler from the indexer for a given namespace and name.
func (s tidbClusterAutoScalerNamespaceLister) Get(name string) (*v1alpha1.TidbClusterAutoScaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tidbclusterautoscaler"), name)
	}
	return obj.(*v1alpha1.TidbClusterAutoScaler), nil
}