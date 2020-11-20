// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type UpstreamWatcher interface {
	// watch namespace-scoped Upstreams
	Watch(namespace string, opts clients.WatchOpts) (<-chan UpstreamList, <-chan error, error)
}

type UpstreamClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Upstream, error)
	Write(resource *Upstream, opts clients.WriteOpts) (*Upstream, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (UpstreamList, error)
	UpstreamWatcher
}

type upstreamClient struct {
	rc clients.ResourceClient
}

func NewUpstreamClient(ctx context.Context, rcFactory factory.ResourceClientFactory) (UpstreamClient, error) {
	return NewUpstreamClientWithToken(ctx, rcFactory, "")
}

func NewUpstreamClientWithToken(ctx context.Context, rcFactory factory.ResourceClientFactory, token string) (UpstreamClient, error) {
	rc, err := rcFactory.NewResourceClient(ctx, factory.NewResourceClientParams{
		ResourceType: &Upstream{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Upstream resource client")
	}
	return NewUpstreamClientWithBase(rc), nil
}

func NewUpstreamClientWithBase(rc clients.ResourceClient) UpstreamClient {
	return &upstreamClient{
		rc: rc,
	}
}

func (client *upstreamClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *upstreamClient) Register() error {
	return client.rc.Register()
}

func (client *upstreamClient) Read(namespace, name string, opts clients.ReadOpts) (*Upstream, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Upstream), nil
}

func (client *upstreamClient) Write(upstream *Upstream, opts clients.WriteOpts) (*Upstream, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(upstream, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Upstream), nil
}

func (client *upstreamClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *upstreamClient) List(namespace string, opts clients.ListOpts) (UpstreamList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToUpstream(resourceList), nil
}

func (client *upstreamClient) Watch(namespace string, opts clients.WatchOpts) (<-chan UpstreamList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	upstreamsChan := make(chan UpstreamList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				upstreamsChan <- convertToUpstream(resourceList)
			case <-opts.Ctx.Done():
				close(upstreamsChan)
				return
			}
		}
	}()
	return upstreamsChan, errs, nil
}

func convertToUpstream(resources resources.ResourceList) UpstreamList {
	var upstreamList UpstreamList
	for _, resource := range resources {
		upstreamList = append(upstreamList, resource.(*Upstream))
	}
	return upstreamList
}
