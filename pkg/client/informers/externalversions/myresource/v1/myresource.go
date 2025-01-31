/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	myresourcev1 "github.com/vlaguduva/kustred/pkg/apis/myresource/v1"
	versioned "github.com/vlaguduva/kustred/pkg/client/clientset/versioned"
	internalinterfaces "github.com/vlaguduva/kustred/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/vlaguduva/kustred/pkg/client/listers/myresource/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MyResourceInformer provides access to a shared informer and lister for
// MyResources.
type MyResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MyResourceLister
}

type myResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMyResourceInformer constructs a new informer for MyResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMyResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMyResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMyResourceInformer constructs a new informer for MyResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMyResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KustredV1().MyResources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KustredV1().MyResources(namespace).Watch(context.TODO(), options)
			},
		},
		&myresourcev1.MyResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *myResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMyResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *myResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&myresourcev1.MyResource{}, f.defaultInformer)
}

func (f *myResourceInformer) Lister() v1.MyResourceLister {
	return v1.NewMyResourceLister(f.Informer().GetIndexer())
}
