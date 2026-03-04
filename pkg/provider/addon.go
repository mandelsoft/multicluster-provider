package provider

import (
	mcpcache "github.com/kcp-dev/multicluster-provider/pkg/cache"
)

func (p *Provider) GetCache() mcpcache.WildcardCache {
	return p.cache
}
