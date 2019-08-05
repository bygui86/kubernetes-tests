/*
Copyright 2019 The KubeDB Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"
)

// ElasticsearchesGetter has a method to return a ElasticsearchInterface.
// A group's client should implement this interface.
type ElasticsearchesGetter interface {
	Elasticsearches(namespace string) ElasticsearchInterface
}

// ElasticsearchInterface has methods to work with Elasticsearch resources.
type ElasticsearchInterface interface {
	Create(*v1alpha1.Elasticsearch) (*v1alpha1.Elasticsearch, error)
	Update(*v1alpha1.Elasticsearch) (*v1alpha1.Elasticsearch, error)
	UpdateStatus(*v1alpha1.Elasticsearch) (*v1alpha1.Elasticsearch, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Elasticsearch, error)
	List(opts v1.ListOptions) (*v1alpha1.ElasticsearchList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Elasticsearch, err error)
	ElasticsearchExpansion
}

// elasticsearches implements ElasticsearchInterface
type elasticsearches struct {
	client rest.Interface
	ns     string
}

// newElasticsearches returns a Elasticsearches
func newElasticsearches(c *KubedbV1alpha1Client, namespace string) *elasticsearches {
	return &elasticsearches{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the elasticsearch, and returns the corresponding elasticsearch object, and an error if there is any.
func (c *elasticsearches) Get(name string, options v1.GetOptions) (result *v1alpha1.Elasticsearch, err error) {
	result = &v1alpha1.Elasticsearch{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearches").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Elasticsearches that match those selectors.
func (c *elasticsearches) List(opts v1.ListOptions) (result *v1alpha1.ElasticsearchList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElasticsearchList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elasticsearches.
func (c *elasticsearches) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elasticsearches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a elasticsearch and creates it.  Returns the server's representation of the elasticsearch, and an error, if there is any.
func (c *elasticsearches) Create(elasticsearch *v1alpha1.Elasticsearch) (result *v1alpha1.Elasticsearch, err error) {
	result = &v1alpha1.Elasticsearch{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elasticsearches").
		Body(elasticsearch).
		Do().
		Into(result)
	return
}

// Update takes the representation of a elasticsearch and updates it. Returns the server's representation of the elasticsearch, and an error, if there is any.
func (c *elasticsearches) Update(elasticsearch *v1alpha1.Elasticsearch) (result *v1alpha1.Elasticsearch, err error) {
	result = &v1alpha1.Elasticsearch{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearches").
		Name(elasticsearch.Name).
		Body(elasticsearch).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *elasticsearches) UpdateStatus(elasticsearch *v1alpha1.Elasticsearch) (result *v1alpha1.Elasticsearch, err error) {
	result = &v1alpha1.Elasticsearch{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elasticsearches").
		Name(elasticsearch.Name).
		SubResource("status").
		Body(elasticsearch).
		Do().
		Into(result)
	return
}

// Delete takes name of the elasticsearch and deletes it. Returns an error if one occurs.
func (c *elasticsearches) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearches").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elasticsearches) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elasticsearches").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched elasticsearch.
func (c *elasticsearches) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Elasticsearch, err error) {
	result = &v1alpha1.Elasticsearch{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elasticsearches").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
