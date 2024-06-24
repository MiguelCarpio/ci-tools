/*


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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	egressfirewallv1 "github.com/openshift/ci-tools/pkg/api/egressfirewall/v1"
	versioned "github.com/openshift/ci-tools/pkg/api/egressfirewall/v1/apis/clientset/versioned"
	internalinterfaces "github.com/openshift/ci-tools/pkg/api/egressfirewall/v1/apis/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/ci-tools/pkg/api/egressfirewall/v1/apis/listers/egressfirewall/v1"
)

// EgressFirewallInformer provides access to a shared informer and lister for
// EgressFirewalls.
type EgressFirewallInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EgressFirewallLister
}

type egressFirewallInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEgressFirewallInformer constructs a new informer for EgressFirewall type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEgressFirewallInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEgressFirewallInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEgressFirewallInformer constructs a new informer for EgressFirewall type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEgressFirewallInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().EgressFirewalls(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().EgressFirewalls(namespace).Watch(context.TODO(), options)
			},
		},
		&egressfirewallv1.EgressFirewall{},
		resyncPeriod,
		indexers,
	)
}

func (f *egressFirewallInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEgressFirewallInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *egressFirewallInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&egressfirewallv1.EgressFirewall{}, f.defaultInformer)
}

func (f *egressFirewallInformer) Lister() v1.EgressFirewallLister {
	return v1.NewEgressFirewallLister(f.Informer().GetIndexer())
}