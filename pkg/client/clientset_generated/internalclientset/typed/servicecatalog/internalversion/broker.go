/*
Copyright 2017 The Kubernetes Authors.

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

package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	api "k8s.io/kubernetes/pkg/api"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	watch "k8s.io/kubernetes/pkg/watch"
)

// BrokersGetter has a method to return a BrokerInterface.
// A group's client should implement this interface.
type BrokersGetter interface {
	Brokers() BrokerInterface
}

// BrokerInterface has methods to work with Broker resources.
type BrokerInterface interface {
	Create(*servicecatalog.Broker) (*servicecatalog.Broker, error)
	Update(*servicecatalog.Broker) (*servicecatalog.Broker, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*servicecatalog.Broker, error)
	List(opts api.ListOptions) (*servicecatalog.BrokerList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *servicecatalog.Broker, err error)
	BrokerExpansion
}

// brokers implements BrokerInterface
type brokers struct {
	client restclient.Interface
}

// newBrokers returns a Brokers
func newBrokers(c *ServicecatalogClient) *brokers {
	return &brokers{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a broker and creates it.  Returns the server's representation of the broker, and an error, if there is any.
func (c *brokers) Create(broker *servicecatalog.Broker) (result *servicecatalog.Broker, err error) {
	result = &servicecatalog.Broker{}
	err = c.client.Post().
		Resource("brokers").
		Body(broker).
		Do().
		Into(result)
	return
}

// Update takes the representation of a broker and updates it. Returns the server's representation of the broker, and an error, if there is any.
func (c *brokers) Update(broker *servicecatalog.Broker) (result *servicecatalog.Broker, err error) {
	result = &servicecatalog.Broker{}
	err = c.client.Put().
		Resource("brokers").
		Name(broker.Name).
		Body(broker).
		Do().
		Into(result)
	return
}

// Delete takes name of the broker and deletes it. Returns an error if one occurs.
func (c *brokers) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Resource("brokers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *brokers) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Resource("brokers").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the broker, and returns the corresponding broker object, and an error if there is any.
func (c *brokers) Get(name string) (result *servicecatalog.Broker, err error) {
	result = &servicecatalog.Broker{}
	err = c.client.Get().
		Resource("brokers").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Brokers that match those selectors.
func (c *brokers) List(opts api.ListOptions) (result *servicecatalog.BrokerList, err error) {
	result = &servicecatalog.BrokerList{}
	err = c.client.Get().
		Resource("brokers").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested brokers.
func (c *brokers) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Resource("brokers").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched broker.
func (c *brokers) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *servicecatalog.Broker, err error) {
	result = &servicecatalog.Broker{}
	err = c.client.Patch(pt).
		Resource("brokers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
