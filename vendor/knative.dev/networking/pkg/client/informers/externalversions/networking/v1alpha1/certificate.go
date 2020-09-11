/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	networkingv1alpha1 "knative.dev/networking/pkg/apis/networking/v1alpha1"
	versioned "knative.dev/networking/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/networking/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "knative.dev/networking/pkg/client/listers/networking/v1alpha1"
)

// CertificateInformer provides access to a shared informer and lister for
// Certificates.
type CertificateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CertificateLister
}

type certificateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCertificateInformer constructs a new informer for Certificate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertificateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateInformer constructs a new informer for Certificate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().Certificates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().Certificates(namespace).Watch(context.TODO(), options)
			},
		},
		&networkingv1alpha1.Certificate{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertificateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *certificateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1alpha1.Certificate{}, f.defaultInformer)
}

func (f *certificateInformer) Lister() v1alpha1.CertificateLister {
	return v1alpha1.NewCertificateLister(f.Informer().GetIndexer())
}
