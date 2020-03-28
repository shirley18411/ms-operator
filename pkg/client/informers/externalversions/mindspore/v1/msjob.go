// Copyright 2020 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	mindsporev1 "gitee.com/mindspore/ms-operator/pkg/apis/mindspore/v1"
	versioned "gitee.com/mindspore/ms-operator/pkg/client/clientset/versioned"
	internalinterfaces "gitee.com/mindspore/ms-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "gitee.com/mindspore/ms-operator/pkg/client/listers/mindspore/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MSJobInformer provides access to a shared informer and lister for
// MSJobs.
type MSJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MSJobLister
}

type mSJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMSJobInformer constructs a new informer for MSJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMSJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMSJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMSJobInformer constructs a new informer for MSJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMSJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().MSJobs(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1().MSJobs(namespace).Watch(options)
			},
		},
		&mindsporev1.MSJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *mSJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMSJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mSJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mindsporev1.MSJob{}, f.defaultInformer)
}

func (f *mSJobInformer) Lister() v1.MSJobLister {
	return v1.NewMSJobLister(f.Informer().GetIndexer())
}
