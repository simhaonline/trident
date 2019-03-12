// Copyright 2019 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TridentSnapshotLister helps list TridentSnapshots.
type TridentSnapshotLister interface {
	// List lists all TridentSnapshots in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentSnapshot, err error)
	// TridentSnapshots returns an object that can list and get TridentSnapshots.
	TridentSnapshots(namespace string) TridentSnapshotNamespaceLister
	TridentSnapshotListerExpansion
}

// tridentSnapshotLister implements the TridentSnapshotLister interface.
type tridentSnapshotLister struct {
	indexer cache.Indexer
}

// NewTridentSnapshotLister returns a new TridentSnapshotLister.
func NewTridentSnapshotLister(indexer cache.Indexer) TridentSnapshotLister {
	return &tridentSnapshotLister{indexer: indexer}
}

// List lists all TridentSnapshots in the indexer.
func (s *tridentSnapshotLister) List(selector labels.Selector) (ret []*v1.TridentSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentSnapshot))
	})
	return ret, err
}

// TridentSnapshots returns an object that can list and get TridentSnapshots.
func (s *tridentSnapshotLister) TridentSnapshots(namespace string) TridentSnapshotNamespaceLister {
	return tridentSnapshotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentSnapshotNamespaceLister helps list and get TridentSnapshots.
type TridentSnapshotNamespaceLister interface {
	// List lists all TridentSnapshots in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentSnapshot, err error)
	// Get retrieves the TridentSnapshot from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentSnapshot, error)
	TridentSnapshotNamespaceListerExpansion
}

// tridentSnapshotNamespaceLister implements the TridentSnapshotNamespaceLister
// interface.
type tridentSnapshotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentSnapshots in the indexer for a given namespace.
func (s tridentSnapshotNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentSnapshot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentSnapshot))
	})
	return ret, err
}

// Get retrieves the TridentSnapshot from the indexer for a given namespace and name.
func (s tridentSnapshotNamespaceLister) Get(name string) (*v1.TridentSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridentsnapshot"), name)
	}
	return obj.(*v1.TridentSnapshot), nil
}
