// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// This file was automatically generated by lister-gen

package v2

import (
	v2 "github.com/northwesternmutual/kanali/pkg/apis/kanali.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApiProxyLister helps list ApiProxies.
type ApiProxyLister interface {
	// List lists all ApiProxies in the indexer.
	List(selector labels.Selector) (ret []*v2.ApiProxy, err error)
	// ApiProxies returns an object that can list and get ApiProxies.
	ApiProxies(namespace string) ApiProxyNamespaceLister
	ApiProxyListerExpansion
}

// apiProxyLister implements the ApiProxyLister interface.
type apiProxyLister struct {
	indexer cache.Indexer
}

// NewApiProxyLister returns a new ApiProxyLister.
func NewApiProxyLister(indexer cache.Indexer) ApiProxyLister {
	return &apiProxyLister{indexer: indexer}
}

// List lists all ApiProxies in the indexer.
func (s *apiProxyLister) List(selector labels.Selector) (ret []*v2.ApiProxy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.ApiProxy))
	})
	return ret, err
}

// ApiProxies returns an object that can list and get ApiProxies.
func (s *apiProxyLister) ApiProxies(namespace string) ApiProxyNamespaceLister {
	return apiProxyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiProxyNamespaceLister helps list and get ApiProxies.
type ApiProxyNamespaceLister interface {
	// List lists all ApiProxies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.ApiProxy, err error)
	// Get retrieves the ApiProxy from the indexer for a given namespace and name.
	Get(name string) (*v2.ApiProxy, error)
	ApiProxyNamespaceListerExpansion
}

// apiProxyNamespaceLister implements the ApiProxyNamespaceLister
// interface.
type apiProxyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiProxies in the indexer for a given namespace.
func (s apiProxyNamespaceLister) List(selector labels.Selector) (ret []*v2.ApiProxy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.ApiProxy))
	})
	return ret, err
}

// Get retrieves the ApiProxy from the indexer for a given namespace and name.
func (s apiProxyNamespaceLister) Get(name string) (*v2.ApiProxy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("apiproxy"), name)
	}
	return obj.(*v2.ApiProxy), nil
}
