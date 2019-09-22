/*
Copyright 2019 BlackRock, Inc.

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

	v1alpha1 "github.com/blackrock/ocibuilder/pkg/apis/ocibuilder/v1alpha1"
	scheme "github.com/blackrock/ocibuilder/pkg/client/ocibuilder/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OCIBuildersGetter has a method to return a OCIBuilderInterface.
// A group's client should implement this interface.
type OCIBuildersGetter interface {
	OCIBuilders(namespace string) OCIBuilderInterface
}

// OCIBuilderInterface has methods to work with OCIBuilder resources.
type OCIBuilderInterface interface {
	Create(*v1alpha1.OCIBuilder) (*v1alpha1.OCIBuilder, error)
	Update(*v1alpha1.OCIBuilder) (*v1alpha1.OCIBuilder, error)
	UpdateStatus(*v1alpha1.OCIBuilder) (*v1alpha1.OCIBuilder, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OCIBuilder, error)
	List(opts v1.ListOptions) (*v1alpha1.OCIBuilderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OCIBuilder, err error)
	OCIBuilderExpansion
}

// oCIBuilders implements OCIBuilderInterface
type oCIBuilders struct {
	client rest.Interface
	ns     string
}

// newOCIBuilders returns a OCIBuilders
func newOCIBuilders(c *BlackrockV1alpha1Client, namespace string) *oCIBuilders {
	return &oCIBuilders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the oCIBuilder, and returns the corresponding oCIBuilder object, and an error if there is any.
func (c *oCIBuilders) Get(name string, options v1.GetOptions) (result *v1alpha1.OCIBuilder, err error) {
	result = &v1alpha1.OCIBuilder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ocibuilders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OCIBuilders that match those selectors.
func (c *oCIBuilders) List(opts v1.ListOptions) (result *v1alpha1.OCIBuilderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OCIBuilderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ocibuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oCIBuilders.
func (c *oCIBuilders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ocibuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a oCIBuilder and creates it.  Returns the server's representation of the oCIBuilder, and an error, if there is any.
func (c *oCIBuilders) Create(oCIBuilder *v1alpha1.OCIBuilder) (result *v1alpha1.OCIBuilder, err error) {
	result = &v1alpha1.OCIBuilder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ocibuilders").
		Body(oCIBuilder).
		Do().
		Into(result)
	return
}

// Update takes the representation of a oCIBuilder and updates it. Returns the server's representation of the oCIBuilder, and an error, if there is any.
func (c *oCIBuilders) Update(oCIBuilder *v1alpha1.OCIBuilder) (result *v1alpha1.OCIBuilder, err error) {
	result = &v1alpha1.OCIBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ocibuilders").
		Name(oCIBuilder.Name).
		Body(oCIBuilder).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *oCIBuilders) UpdateStatus(oCIBuilder *v1alpha1.OCIBuilder) (result *v1alpha1.OCIBuilder, err error) {
	result = &v1alpha1.OCIBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ocibuilders").
		Name(oCIBuilder.Name).
		SubResource("status").
		Body(oCIBuilder).
		Do().
		Into(result)
	return
}

// Delete takes name of the oCIBuilder and deletes it. Returns an error if one occurs.
func (c *oCIBuilders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ocibuilders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oCIBuilders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ocibuilders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched oCIBuilder.
func (c *oCIBuilders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OCIBuilder, err error) {
	result = &v1alpha1.OCIBuilder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ocibuilders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
