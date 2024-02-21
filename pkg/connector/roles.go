package connector

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-celigo/pkg/celigo"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	grant "github.com/conductorone/baton-sdk/pkg/types/grant"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
)

const (
	MonitorAllRole    = "monitor"
	ManageAllRole     = "manage"
	AdministratorRole = "administrator"

	assignedEntitlement = "assigned"
)

var roles = []string{MonitorAllRole, ManageAllRole, AdministratorRole}

type roleBuilder struct {
	resourceType *v2.ResourceType
	client       *celigo.Client
}

func (r *roleBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return roleResourceType
}

func newRoleResource(ctx context.Context, role string) (*v2.Resource, error) {
	profile := map[string]interface{}{
		"name": role,
	}

	roleTraitOptions := []rs.RoleTraitOption{
		rs.WithRoleProfile(profile),
	}

	resource, err := rs.NewRoleResource(role, roleResourceType, role, roleTraitOptions)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (r *roleBuilder) List(ctx context.Context, _ *v2.ResourceId, _ *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	var resources []*v2.Resource
	for _, role := range roles {
		resource, err := newRoleResource(ctx, role)
		if err != nil {
			return nil, "", nil, err
		}

		resources = append(resources, resource)
	}

	return resources, "", nil, nil
}

func (o *roleBuilder) Entitlements(ctx context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	assigmentOptions := []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDescription(fmt.Sprintf("Assigned to %s role", resource.DisplayName)),
		ent.WithDisplayName(fmt.Sprintf("%s role %s", resource.DisplayName, assignedEntitlement)),
	}
	rv = append(rv, ent.NewAssignmentEntitlement(resource, assignedEntitlement, assigmentOptions...))

	return rv, "", nil, nil
}

func (o *roleBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	bag, nextPageLink, err := parsePageToken(pToken.Token, &v2.ResourceId{ResourceType: o.resourceType.Id})
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get next page link")
	}

	response, nextPageLink, _, err := o.client.ListUsers(ctx, nextPageLink)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to list users")
	}

	var rv []*v2.Grant
	for _, user := range response {
		if user.AccessLevel != resource.Id.Resource {
			continue
		}

		user := user
		userResource, err := userResource(ctx, &user)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create user resource")
		}

		rv = append(rv, grant.NewGrant(resource, assignedEntitlement, userResource.Id))
	}

	if nextPageLink != "" {
		nextPageLink, err = handleNextPage(bag, nextPageLink)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create next page cursor")
		}
	}

	return rv, nextPageLink, nil, nil
}

func newRoleBuilder(client *celigo.Client) *roleBuilder {
	return &roleBuilder{
		resourceType: roleResourceType,
		client:       client,
	}
}
