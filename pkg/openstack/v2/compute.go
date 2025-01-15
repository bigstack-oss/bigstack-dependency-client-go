package openstack

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/quotasets"
)

func (h *Helper) UpdateComputeQuotas(projectId string, opts quotasets.UpdateOpts) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return quotasets.Update(ctx, h.Compute, projectId, opts).Err
}
