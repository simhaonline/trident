// Copyright 2019 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/netapp/trident/persistent_store/crd/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// TridentBackends returns a TridentBackendInformer.
	TridentBackends() TridentBackendInformer
	// TridentNodes returns a TridentNodeInformer.
	TridentNodes() TridentNodeInformer
	// TridentSnapshots returns a TridentSnapshotInformer.
	TridentSnapshots() TridentSnapshotInformer
	// TridentStorageClasses returns a TridentStorageClassInformer.
	TridentStorageClasses() TridentStorageClassInformer
	// TridentTransactions returns a TridentTransactionInformer.
	TridentTransactions() TridentTransactionInformer
	// TridentVersions returns a TridentVersionInformer.
	TridentVersions() TridentVersionInformer
	// TridentVolumes returns a TridentVolumeInformer.
	TridentVolumes() TridentVolumeInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// TridentBackends returns a TridentBackendInformer.
func (v *version) TridentBackends() TridentBackendInformer {
	return &tridentBackendInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentNodes returns a TridentNodeInformer.
func (v *version) TridentNodes() TridentNodeInformer {
	return &tridentNodeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentSnapshots returns a TridentSnapshotInformer.
func (v *version) TridentSnapshots() TridentSnapshotInformer {
	return &tridentSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentStorageClasses returns a TridentStorageClassInformer.
func (v *version) TridentStorageClasses() TridentStorageClassInformer {
	return &tridentStorageClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentTransactions returns a TridentTransactionInformer.
func (v *version) TridentTransactions() TridentTransactionInformer {
	return &tridentTransactionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentVersions returns a TridentVersionInformer.
func (v *version) TridentVersions() TridentVersionInformer {
	return &tridentVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TridentVolumes returns a TridentVolumeInformer.
func (v *version) TridentVolumes() TridentVolumeInformer {
	return &tridentVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
