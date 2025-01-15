package openstack

import (
	"context"
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
)

func (h *Helper) AddRole(roleId string, opts roles.AssignOpts) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return roles.Assign(
		ctx,
		h.Identity,
		roleId,
		opts,
	).ExtractErr()
}

func (h *Helper) GetRoleByName(name string) (*roles.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pages, err := roles.List(h.Identity, roles.ListOpts{Name: name}).AllPages(ctx)
	if err != nil {
		return nil, err
	}

	roles, err := roles.ExtractRoles(pages)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		if role.Name == name {
			return &role, nil
		}
	}

	return nil, fmt.Errorf("role %s not found", name)
}
