package openstack

import (
	"context"
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
)

func (h *Helper) ListProjects(opts *projects.ListOpts) ([]projects.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pages, err := projects.List(h.Identity, opts).AllPages(ctx)
	if err != nil {
		return nil, err
	}

	return projects.ExtractProjects(pages)
}

func (h *Helper) GetProjectIdByName(name string) (string, error) {
	projects, err := h.ListProjects(&projects.ListOpts{Name: name})
	if err != nil {
		return "", err
	}

	projectId := ""
	for _, project := range projects {
		if project.Name == name {
			projectId = project.ID
			break
		}
	}
	if projectId == "" {
		return "", fmt.Errorf("project %s not found", name)
	}

	return projectId, nil
}

func (h *Helper) CreateProject(name string) (*projects.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	true := true
	return projects.Create(
		ctx,
		h.Identity,
		projects.CreateOpts{
			Name:    name,
			Enabled: &true,
		},
	).Extract()
}
