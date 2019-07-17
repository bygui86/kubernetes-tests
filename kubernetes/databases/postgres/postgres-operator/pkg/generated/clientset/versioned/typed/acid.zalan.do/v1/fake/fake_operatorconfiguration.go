/*
Copyright 2019 Compose, Zalando SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	acidzalandov1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeOperatorConfigurations implements OperatorConfigurationInterface
type FakeOperatorConfigurations struct {
	Fake *FakeAcidV1
	ns   string
}

var operatorconfigurationsResource = schema.GroupVersionResource{Group: "acid.zalan.do", Version: "v1", Resource: "operatorconfigurations"}

var operatorconfigurationsKind = schema.GroupVersionKind{Group: "acid.zalan.do", Version: "v1", Kind: "OperatorConfiguration"}

// Get takes name of the operatorConfiguration, and returns the corresponding operatorConfiguration object, and an error if there is any.
func (c *FakeOperatorConfigurations) Get(name string, options v1.GetOptions) (result *acidzalandov1.OperatorConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operatorconfigurationsResource, c.ns, name), &acidzalandov1.OperatorConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidzalandov1.OperatorConfiguration), err
}