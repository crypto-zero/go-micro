// Package static provides a static resolver which returns the name/ip passed in without any change
package static

import (
	"context"
	"github.com/crypto-zero/go-micro/v2/client/selector"
	"github.com/crypto-zero/go-micro/v2/registry"
)

type staticAddressOption struct{}

// staticSelector is a static selector
type staticSelector struct {
	opts selector.Options
}

func (s *staticSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&s.opts)
	}
	return nil
}

func (s *staticSelector) Options() selector.Options {
	return s.opts
}

func (s *staticSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	address := service
	if addressValue := s.opts.Context.Value(staticAddressOption{}); addressValue != nil {
		address = addressValue.(string)
	}
	return func() (*registry.Node, error) {
		return &registry.Node{
			Id:      service,
			Address: address,
		}, nil
	}, nil
}

func (s *staticSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (s *staticSelector) Reset(service string) {
	return
}

func (s *staticSelector) Close() error {
	return nil
}

func (s *staticSelector) String() string {
	return "static"
}

func WithStaticAddress(address string) selector.Option {
	return func(o *selector.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, staticAddressOption{}, address)
	}
}

func NewSelector(opts ...selector.Option) selector.Selector {
	var options selector.Options
	for _, o := range opts {
		o(&options)
	}
	return &staticSelector{
		opts: options,
	}
}
