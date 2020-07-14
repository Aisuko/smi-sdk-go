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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	scheme "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/access/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TrafficTargetsGetter has a method to return a TrafficTargetInterface.
// A group's client should implement this interface.
type TrafficTargetsGetter interface {
	TrafficTargets(namespace string) TrafficTargetInterface
}

// TrafficTargetInterface has methods to work with TrafficTarget resources.
type TrafficTargetInterface interface {
	Create(ctx context.Context, trafficTarget *v1alpha2.TrafficTarget, opts v1.CreateOptions) (*v1alpha2.TrafficTarget, error)
	Update(ctx context.Context, trafficTarget *v1alpha2.TrafficTarget, opts v1.UpdateOptions) (*v1alpha2.TrafficTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.TrafficTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.TrafficTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TrafficTarget, err error)
	TrafficTargetExpansion
}

// trafficTargets implements TrafficTargetInterface
type trafficTargets struct {
	client rest.Interface
	ns     string
}

// newTrafficTargets returns a TrafficTargets
func newTrafficTargets(c *AccessV1alpha2Client, namespace string) *trafficTargets {
	return &trafficTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the trafficTarget, and returns the corresponding trafficTarget object, and an error if there is any.
func (c *trafficTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.TrafficTarget, err error) {
	result = &v1alpha2.TrafficTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("traffictargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TrafficTargets that match those selectors.
func (c *trafficTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.TrafficTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.TrafficTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("traffictargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested trafficTargets.
func (c *trafficTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("traffictargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a trafficTarget and creates it.  Returns the server's representation of the trafficTarget, and an error, if there is any.
func (c *trafficTargets) Create(ctx context.Context, trafficTarget *v1alpha2.TrafficTarget, opts v1.CreateOptions) (result *v1alpha2.TrafficTarget, err error) {
	result = &v1alpha2.TrafficTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("traffictargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trafficTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a trafficTarget and updates it. Returns the server's representation of the trafficTarget, and an error, if there is any.
func (c *trafficTargets) Update(ctx context.Context, trafficTarget *v1alpha2.TrafficTarget, opts v1.UpdateOptions) (result *v1alpha2.TrafficTarget, err error) {
	result = &v1alpha2.TrafficTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("traffictargets").
		Name(trafficTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trafficTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the trafficTarget and deletes it. Returns an error if one occurs.
func (c *trafficTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("traffictargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *trafficTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("traffictargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched trafficTarget.
func (c *trafficTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TrafficTarget, err error) {
	result = &v1alpha2.TrafficTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("traffictargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
