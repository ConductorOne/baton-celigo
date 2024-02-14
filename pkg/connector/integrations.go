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
	MonitorItegrationEntitlement = "monitor"
	ManageIntegrationEntitlement = "manage"
)

type integrationsBuilder struct {
	resourceType *v2.ResourceType
	client       *celigo.Client
}

func (o *integrationsBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return integrationsResourceType
}

func integrationResource(ctx context.Context, integration *celigo.Integration) (*v2.Resource, error) {
	profile := map[string]interface{}{
		"name":        integration.Name,
		"id":          integration.Id,
		"description": integration.Description,
	}

	integrationTraitOptions := []rs.AppTraitOption{
		rs.WithAppProfile(profile),
	}
	resource, err := rs.NewAppResource(integration.Name, integrationsResourceType, integration.Id, integrationTraitOptions)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

// List returns all the users from the database as resource objects.
// Users include a UserTrait because they are the 'shape' of a standard user.
func (o *integrationsBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	bag, nextPageLink, err := parsePageToken(pToken.Token, &v2.ResourceId{ResourceType: o.resourceType.Id})
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get next page link")
	}

	response, nextPageLink, _, err := o.client.ListIntegrations(ctx, nextPageLink)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to list integrations")
	}

	var resources []*v2.Resource
	for _, integration := range response {
		integration := integration
		resource, err := integrationResource(ctx, &integration)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create integration resource")
		}

		resources = append(resources, resource)
	}

	if nextPageLink == "" {
		nextPageLink, err = handleNextPage(bag, nextPageLink)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create next page cursor")
		}
	}

	return resources, nextPageLink, nil, nil
}

// Entitlements always returns an empty slice for users.
func (o *integrationsBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	assigmentOptions := []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDescription(fmt.Sprintf("has %s access level", resource.DisplayName)),
		ent.WithDisplayName(fmt.Sprintf("%s access level %s", resource.DisplayName, ManageIntegrationEntitlement)),
	}
	rv = append(rv, ent.NewAssignmentEntitlement(resource, ManageIntegrationEntitlement, assigmentOptions...))

	assigmentOptions = []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDescription(fmt.Sprintf("has %s access level", resource.DisplayName)),
		ent.WithDisplayName(fmt.Sprintf("%s access level %s", resource.DisplayName, MonitorItegrationEntitlement)),
	}
	rv = append(rv, ent.NewAssignmentEntitlement(resource, MonitorItegrationEntitlement, assigmentOptions...))

	return rv, "", nil, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (o *integrationsBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	bag, nextPageLink, err := parsePageToken(pToken.Token, &v2.ResourceId{ResourceType: o.resourceType.Id})
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get next page link")
	}

	response, nextPageLink, _, err := o.client.ListIntegrationsUsers(ctx, resource.Id.Resource, nextPageLink)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to list users")
	}

	var rv []*v2.Grant
	for _, user := range response {
		user := user
		userResource, err := userResource(ctx, &user)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create user resource")
		}

		rv = append(rv, grant.NewGrant(resource, user.AccessLevel, userResource.Id))
	}

	if nextPageLink != "" {
		nextPageLink, err = handleNextPage(bag, nextPageLink)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create next page cursor")
		}
	}

	return rv, nextPageLink, nil, nil
}

func newIntegrationsBuilder(client *celigo.Client) *integrationsBuilder {
	return &integrationsBuilder{
		resourceType: userResourceType,
		client:       client,
	}
}
