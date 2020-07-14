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

package fake

import (
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/specs/clientset/versioned/typed/specs/v1alpha3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSpecsV1alpha3 struct {
	*testing.Fake
}

func (c *FakeSpecsV1alpha3) HTTPRouteGroups(namespace string) v1alpha3.HTTPRouteGroupInterface {
	return &FakeHTTPRouteGroups{c, namespace}
}

func (c *FakeSpecsV1alpha3) TCPRoutes(namespace string) v1alpha3.TCPRouteInterface {
	return &FakeTCPRoutes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSpecsV1alpha3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
