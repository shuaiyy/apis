/*
Copyright The Volcano Authors.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	schedulingv1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
)

// QueueLister helps list Queues.
// All objects returned here must be treated as read-only.
type QueueLister interface {
	// List lists all Queues in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*schedulingv1beta1.Queue, err error)
	// Get retrieves the Queue from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*schedulingv1beta1.Queue, error)
	QueueListerExpansion
}

// queueLister implements the QueueLister interface.
type queueLister struct {
	listers.ResourceIndexer[*schedulingv1beta1.Queue]
}

// NewQueueLister returns a new QueueLister.
func NewQueueLister(indexer cache.Indexer) QueueLister {
	return &queueLister{listers.New[*schedulingv1beta1.Queue](indexer, schedulingv1beta1.Resource("queue"))}
}
