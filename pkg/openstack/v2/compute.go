package openstack

import (
	"context"
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/hypervisors"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/quotasets"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
)

func (h *Helper) UpdateComputeQuotas(projectId string, opts quotasets.UpdateOpts) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return quotasets.Update(ctx, h.Compute, projectId, opts).Err
}

func (h *Helper) ListServers(opts servers.ListOpts) ([]servers.Server, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pages, err := servers.List(h.Compute, opts).AllPages(ctx)
	if err != nil {
		return nil, err
	}

	return servers.ExtractServers(pages)
}

func (h *Helper) GetHypervisorStatistics() (*hypervisors.Statistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return hypervisors.GetStatistics(ctx, h.Compute).Extract()
}

func (h *Helper) ListHypervisors(opts hypervisors.ListOpts) ([]hypervisors.Hypervisor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pages, err := hypervisors.List(h.Compute, hypervisors.ListOpts{}).AllPages(ctx)
	if err != nil {
		return nil, err
	}

	return hypervisors.ExtractHypervisors(pages)
}

func (h *Helper) GetHypervisorByHostname(hostname string) (*hypervisors.Hypervisor, error) {
	hps, err := h.ListHypervisors(hypervisors.ListOpts{})
	if err != nil {
		return nil, err
	}

	err = fmt.Errorf("hypervisor with hostname %s not found", hostname)
	if len(hps) == 0 {
		return nil, err
	}

	for _, hypervisor := range hps {
		if hypervisor.HypervisorHostname == hostname {
			return &hypervisor, nil
		}
	}

	return nil, err
}

func (h *Helper) GetHypervisorUpTime(id string) (*hypervisors.Uptime, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return hypervisors.GetUptime(ctx, h.Compute, id).Extract()
}
