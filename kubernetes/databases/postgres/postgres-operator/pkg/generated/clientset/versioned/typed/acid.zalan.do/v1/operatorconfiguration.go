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

package v1

import (
	acidzalandov1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	scheme "github.com/zalando/postgres-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// OperatorConfigurationsGetter has a method to return a OperatorConfigurationInterface.
// A group's client should implement this interface.
type OperatorConfigurationsGetter interface {
	OperatorConfigurations(namespace string) OperatorConfigurationInterface
}

// OperatorConfigurationInterface has methods to work with OperatorConfiguration resources.
type OperatorConfigurationInterface interface {
	Get(name string, options v1.GetOptions) (*acidzalandov1.OperatorConfiguration, error)
	OperatorConfigurationExpansion
}

// operatorConfigurations implements OperatorConfigurationInterface
type operatorConfigurations struct {
	client rest.Interface
	ns     string
}

// newOperatorConfigurations returns a OperatorConfigurations
func newOperatorConfigurations(c *AcidV1Client, namespace string) *operatorConfigurations {
	return &operatorConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the operatorConfiguration, and returns the corresponding operatorConfiguration object, and an error if there is any.
func (c *operatorConfigurations) Get(name string, options v1.GetOptions) (result *acidzalandov1.OperatorConfiguration, err error) {
	result = &acidzalandov1.OperatorConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}
