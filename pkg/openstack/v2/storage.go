package openstack

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/quotasets"
)

func (h *Helper) UpdateStorageQuotas(projectId string, opts quotasets.UpdateOpts) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return quotasets.Update(ctx, h.Storage, projectId, opts).Err
}
