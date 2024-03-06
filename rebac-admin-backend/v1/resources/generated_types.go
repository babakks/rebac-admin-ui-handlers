// Package resources provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package resources

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for CapabilityMethods.
const (
	DELETE CapabilityMethods = "DELETE"
	GET    CapabilityMethods = "GET"
	PATCH  CapabilityMethods = "PATCH"
	POST   CapabilityMethods = "POST"
)

// Defines values for GroupEntitlementsPatchItemOp.
const (
	GroupEntitlementsPatchItemOpAdd    GroupEntitlementsPatchItemOp = "add"
	GroupEntitlementsPatchItemOpRemove GroupEntitlementsPatchItemOp = "remove"
)

// Defines values for GroupIdentitiesPatchItemOp.
const (
	GroupIdentitiesPatchItemOpAdd    GroupIdentitiesPatchItemOp = "add"
	GroupIdentitiesPatchItemOpRemove GroupIdentitiesPatchItemOp = "remove"
)

// Defines values for GroupRolesPatchItemOp.
const (
	GroupRolesPatchItemOpAdd    GroupRolesPatchItemOp = "add"
	GroupRolesPatchItemOpRemove GroupRolesPatchItemOp = "remove"
)

// Defines values for IdentityEntitlementsPatchItemOp.
const (
	IdentityEntitlementsPatchItemOpAdd    IdentityEntitlementsPatchItemOp = "add"
	IdentityEntitlementsPatchItemOpRemove IdentityEntitlementsPatchItemOp = "remove"
)

// Defines values for IdentityGroupsPatchItemOp.
const (
	IdentityGroupsPatchItemOpAdd    IdentityGroupsPatchItemOp = "add"
	IdentityGroupsPatchItemOpRemove IdentityGroupsPatchItemOp = "remove"
)

// Defines values for IdentityProviderSyncMode.
const (
	Import IdentityProviderSyncMode = "import"
)

// Defines values for IdentityRolesPatchItemOp.
const (
	IdentityRolesPatchItemOpAdd    IdentityRolesPatchItemOp = "add"
	IdentityRolesPatchItemOpRemove IdentityRolesPatchItemOp = "remove"
)

// Defines values for RoleEntitlementsPatchItemOp.
const (
	Add    RoleEntitlementsPatchItemOp = "add"
	Remove RoleEntitlementsPatchItemOp = "remove"
)

// AuthenticationProvider defines model for AuthenticationProvider.
type AuthenticationProvider struct {
	Id   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// AuthenticationProviders defines model for AuthenticationProviders.
type AuthenticationProviders struct {
	Data []AuthenticationProvider `json:"data"`
}

// Capabilities defines model for Capabilities.
type Capabilities struct {
	Data []Capability `json:"data"`
}

// Capability defines model for Capability.
type Capability struct {
	Endpoint string              `json:"endpoint"`
	Methods  []CapabilityMethods `json:"methods"`
}

// CapabilityMethods defines model for Capability.Methods.
type CapabilityMethods string

// Entity defines model for Entity.
type Entity = map[string]interface{}

// EntityEntitlement defines model for EntityEntitlement.
type EntityEntitlement struct {
	EntitlementType string `json:"entitlement_type"`
	EntityName      string `json:"entity_name"`
	EntityType      string `json:"entity_type"`
}

// EntityEntitlementItem defines model for EntityEntitlementItem.
type EntityEntitlementItem struct {
	Entitlement EntityEntitlement `json:"entitlement"`
}

// GetAuthenticationProvidersResponse defines model for GetAuthenticationProvidersResponse.
type GetAuthenticationProvidersResponse struct {
	Links   ResponseLinks            `json:"_links"`
	Meta    ResponseMeta             `json:"_meta"`
	Data    []AuthenticationProvider `json:"data"`
	Message string                   `json:"message"`
	Status  int                      `json:"status"`
}

// GetCapabilitiesResponse defines model for GetCapabilitiesResponse.
type GetCapabilitiesResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Capability  `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetEntitlementsResponse defines model for GetEntitlementsResponse.
type GetEntitlementsResponse struct {
	Links   ResponseLinks       `json:"_links"`
	Meta    ResponseMeta        `json:"_meta"`
	Data    []EntityEntitlement `json:"data"`
	Message string              `json:"message"`
	Status  int                 `json:"status"`
}

// GetGroupEntitlementsResponse defines model for GetGroupEntitlementsResponse.
type GetGroupEntitlementsResponse struct {
	Links   ResponseLinks       `json:"_links"`
	Meta    ResponseMeta        `json:"_meta"`
	Data    []EntityEntitlement `json:"data"`
	Message string              `json:"message"`
	Status  int                 `json:"status"`
}

// GetGroupIdentitiesResponse defines model for GetGroupIdentitiesResponse.
type GetGroupIdentitiesResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Identity    `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetGroupRolesResponse defines model for GetGroupRolesResponse.
type GetGroupRolesResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Role        `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetGroupsResponse defines model for GetGroupsResponse.
type GetGroupsResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Group       `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetIdentitiesResponse defines model for GetIdentitiesResponse.
type GetIdentitiesResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Identity    `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetIdentityEntitlementsResponse defines model for GetIdentityEntitlementsResponse.
type GetIdentityEntitlementsResponse struct {
	Links   ResponseLinks       `json:"_links"`
	Meta    ResponseMeta        `json:"_meta"`
	Data    []EntityEntitlement `json:"data"`
	Message string              `json:"message"`
	Status  int                 `json:"status"`
}

// GetIdentityGroupsResponse defines model for GetIdentityGroupsResponse.
type GetIdentityGroupsResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Group       `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetIdentityProvidersResponse defines model for GetIdentityProvidersResponse.
type GetIdentityProvidersResponse struct {
	Links   ResponseLinks      `json:"_links"`
	Meta    ResponseMeta       `json:"_meta"`
	Data    []IdentityProvider `json:"data"`
	Message string             `json:"message"`
	Status  int                `json:"status"`
}

// GetIdentityRolesResponse defines model for GetIdentityRolesResponse.
type GetIdentityRolesResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Role        `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetResourcesResponse defines model for GetResourcesResponse.
type GetResourcesResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Resource    `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// GetRoleEntitlementsResponse defines model for GetRoleEntitlementsResponse.
type GetRoleEntitlementsResponse struct {
	Links   ResponseLinks       `json:"_links"`
	Meta    ResponseMeta        `json:"_meta"`
	Data    []EntityEntitlement `json:"data"`
	Message string              `json:"message"`
	Status  int                 `json:"status"`
}

// GetRolesResponse defines model for GetRolesResponse.
type GetRolesResponse struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Data    []Role        `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// Group defines model for Group.
type Group struct {
	Id   *string `json:"id,omitempty"`
	Name string  `json:"name"`
}

// GroupEntitlementsPatchItem defines model for GroupEntitlementsPatchItem.
type GroupEntitlementsPatchItem struct {
	Entitlement EntityEntitlement            `json:"entitlement"`
	Op          GroupEntitlementsPatchItemOp `json:"op"`
}

// GroupEntitlementsPatchItemOp defines model for GroupEntitlementsPatchItem.Op.
type GroupEntitlementsPatchItemOp string

// GroupEntitlementsPatchRequestBody defines model for GroupEntitlementsPatchRequestBody.
type GroupEntitlementsPatchRequestBody struct {
	Patches []GroupEntitlementsPatchItem `json:"patches"`
}

// GroupIdentitiesPatchItem defines model for GroupIdentitiesPatchItem.
type GroupIdentitiesPatchItem struct {
	Identity string                     `json:"identity"`
	Op       GroupIdentitiesPatchItemOp `json:"op"`
}

// GroupIdentitiesPatchItemOp defines model for GroupIdentitiesPatchItem.Op.
type GroupIdentitiesPatchItemOp string

// GroupIdentitiesPatchRequestBody defines model for GroupIdentitiesPatchRequestBody.
type GroupIdentitiesPatchRequestBody struct {
	Patches []GroupIdentitiesPatchItem `json:"patches"`
}

// GroupRolesPatchItem defines model for GroupRolesPatchItem.
type GroupRolesPatchItem struct {
	Op   GroupRolesPatchItemOp `json:"op"`
	Role string                `json:"role"`
}

// GroupRolesPatchItemOp defines model for GroupRolesPatchItem.Op.
type GroupRolesPatchItemOp string

// GroupRolesPatchRequestBody defines model for GroupRolesPatchRequestBody.
type GroupRolesPatchRequestBody struct {
	Patches []GroupRolesPatchItem `json:"patches"`
}

// Groups defines model for Groups.
type Groups struct {
	Data []Group `json:"data"`
}

// Identities defines model for Identities.
type Identities struct {
	Data []Identity `json:"data"`
}

// Identity defines model for Identity.
type Identity struct {
	AddedBy     string  `json:"addedBy"`
	Certificate *string `json:"certificate,omitempty"`
	Email       string  `json:"email"`
	FirstName   *string `json:"firstName,omitempty"`
	Groups      *int    `json:"groups,omitempty"`
	Id          *string `json:"id,omitempty"`
	Joined      *string `json:"joined,omitempty"`
	LastLogin   *string `json:"lastLogin,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	Permissions *int    `json:"permissions,omitempty"`
	Roles       *int    `json:"roles,omitempty"`
	Source      string  `json:"source"`
}

// IdentityEntitlementsPatchItem defines model for IdentityEntitlementsPatchItem.
type IdentityEntitlementsPatchItem struct {
	Entitlement EntityEntitlement               `json:"entitlement"`
	Op          IdentityEntitlementsPatchItemOp `json:"op"`
}

// IdentityEntitlementsPatchItemOp defines model for IdentityEntitlementsPatchItem.Op.
type IdentityEntitlementsPatchItemOp string

// IdentityEntitlementsPatchRequestBody defines model for IdentityEntitlementsPatchRequestBody.
type IdentityEntitlementsPatchRequestBody struct {
	Patches []IdentityEntitlementsPatchItem `json:"patches"`
}

// IdentityGroupsPatchItem defines model for IdentityGroupsPatchItem.
type IdentityGroupsPatchItem struct {
	Group string                    `json:"group"`
	Op    IdentityGroupsPatchItemOp `json:"op"`
}

// IdentityGroupsPatchItemOp defines model for IdentityGroupsPatchItem.Op.
type IdentityGroupsPatchItemOp string

// IdentityGroupsPatchRequestBody defines model for IdentityGroupsPatchRequestBody.
type IdentityGroupsPatchRequestBody struct {
	Patches []IdentityGroupsPatchItem `json:"patches"`
}

// IdentityProvider defines model for IdentityProvider.
type IdentityProvider struct {
	AcceptsPromptNone   *bool                     `json:"acceptsPromptNone,omitempty"`
	AccountLinkingOnly  *bool                     `json:"accountLinkingOnly,omitempty"`
	ClientID            *string                   `json:"clientID,omitempty"`
	ClientSecret        *string                   `json:"clientSecret,omitempty"`
	DisableIdentityInfo *bool                     `json:"disableIdentityInfo,omitempty"`
	Enabled             *bool                     `json:"enabled,omitempty"`
	Id                  *string                   `json:"id,omitempty"`
	IdentityCount       *int                      `json:"identityCount,omitempty"`
	Name                *string                   `json:"name,omitempty"`
	RedirectUrl         *string                   `json:"redirectUrl,omitempty"`
	StoreTokens         *bool                     `json:"storeTokens,omitempty"`
	StoreTokensReadable *bool                     `json:"storeTokensReadable,omitempty"`
	SyncMode            *IdentityProviderSyncMode `json:"syncMode,omitempty"`
	TrustEmail          *bool                     `json:"trustEmail,omitempty"`
}

// IdentityProviderSyncMode defines model for IdentityProvider.SyncMode.
type IdentityProviderSyncMode string

// IdentityProviders defines model for IdentityProviders.
type IdentityProviders struct {
	Data []IdentityProvider `json:"data"`
}

// IdentityRolesPatchItem defines model for IdentityRolesPatchItem.
type IdentityRolesPatchItem struct {
	Op   IdentityRolesPatchItemOp `json:"op"`
	Role string                   `json:"role"`
}

// IdentityRolesPatchItemOp defines model for IdentityRolesPatchItem.Op.
type IdentityRolesPatchItemOp string

// IdentityRolesPatchRequestBody defines model for IdentityRolesPatchRequestBody.
type IdentityRolesPatchRequestBody struct {
	Patches []IdentityRolesPatchItem `json:"patches"`
}

// Resource defines model for Resource.
type Resource struct {
	Entity Entity    `json:"entity"`
	Id     string    `json:"id"`
	Name   string    `json:"name"`
	Parent *Resource `json:"parent,omitempty"`
}

// Resources defines model for Resources.
type Resources struct {
	Data []Resource `json:"data"`
}

// Response defines model for Response.
type Response struct {
	Links   ResponseLinks `json:"_links"`
	Meta    ResponseMeta  `json:"_meta"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

// ResponseLinks defines model for ResponseLinks.
type ResponseLinks struct {
	Next ResponseLinksNext `json:"next"`
}

// ResponseLinksNext defines model for ResponseLinksNext.
type ResponseLinksNext struct {
	Href string `json:"href"`
}

// ResponseMeta defines model for ResponseMeta.
type ResponseMeta struct {
	Page      *int    `json:"page,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	Size      int     `json:"size"`
	Total     *int    `json:"total,omitempty"`
}

// Role defines model for Role.
type Role struct {
	Entitlements *[]RoleEntitlement `json:"entitlements,omitempty"`
	Id           *string            `json:"id,omitempty"`
	Name         string             `json:"name"`
}

// RoleEntitlement defines model for RoleEntitlement.
type RoleEntitlement struct {
	Entitlement *string `json:"entitlement,omitempty"`
	Entity      *Entity `json:"entity,omitempty"`
	Resource    *string `json:"resource,omitempty"`
}

// RoleEntitlementsPatchItem defines model for RoleEntitlementsPatchItem.
type RoleEntitlementsPatchItem struct {
	Entitlement EntityEntitlement           `json:"entitlement"`
	Op          RoleEntitlementsPatchItemOp `json:"op"`
}

// RoleEntitlementsPatchItemOp defines model for RoleEntitlementsPatchItem.Op.
type RoleEntitlementsPatchItemOp string

// RoleEntitlementsPatchRequestBody defines model for RoleEntitlementsPatchRequestBody.
type RoleEntitlementsPatchRequestBody struct {
	Patches []RoleEntitlementsPatchItem `json:"patches"`
}

// Roles defines model for Roles.
type Roles struct {
	Data []Role `json:"data"`
}

// FilterParam defines model for FilterParam.
type FilterParam = string

// PaginationNextToken defines model for PaginationNextToken.
type PaginationNextToken = string

// PaginationPage defines model for PaginationPage.
type PaginationPage = int

// PaginationSize defines model for PaginationSize.
type PaginationSize = int

// BadRequest defines model for BadRequest.
type BadRequest = Response

// NotFound defines model for NotFound.
type NotFound = Response

// Unauthorized defines model for Unauthorized.
type Unauthorized = Response

// Default defines model for default.
type Default = Response

// GetAuthenticationParams defines parameters for GetAuthentication.
type GetAuthenticationParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetAuthenticationProvidersParams defines parameters for GetAuthenticationProviders.
type GetAuthenticationProvidersParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetEntitlementsParams defines parameters for GetEntitlements.
type GetEntitlementsParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`

	// Filter A string to filter results by
	Filter *FilterParam `form:"filter,omitempty" json:"filter,omitempty"`
}

// GetGroupsParams defines parameters for GetGroups.
type GetGroupsParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`

	// Filter A string to filter results by
	Filter *FilterParam `form:"filter,omitempty" json:"filter,omitempty"`
}

// GetGroupsItemEntitlementsParams defines parameters for GetGroupsItemEntitlements.
type GetGroupsItemEntitlementsParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetGroupsItemIdentitiesParams defines parameters for GetGroupsItemIdentities.
type GetGroupsItemIdentitiesParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetGroupsItemRolesParams defines parameters for GetGroupsItemRoles.
type GetGroupsItemRolesParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetIdentitiesParams defines parameters for GetIdentities.
type GetIdentitiesParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`

	// Filter A string to filter results by
	Filter *FilterParam `form:"filter,omitempty" json:"filter,omitempty"`
}

// GetIdentitiesItemEntitlementsParams defines parameters for GetIdentitiesItemEntitlements.
type GetIdentitiesItemEntitlementsParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetIdentitiesItemGroupsParams defines parameters for GetIdentitiesItemGroups.
type GetIdentitiesItemGroupsParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetIdentitiesItemRolesParams defines parameters for GetIdentitiesItemRoles.
type GetIdentitiesItemRolesParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// GetResourcesParams defines parameters for GetResources.
type GetResourcesParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken  *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
	EntityType *string              `form:"entityType,omitempty" json:"entityType,omitempty"`
}

// GetRolesParams defines parameters for GetRoles.
type GetRolesParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`

	// Filter A string to filter results by
	Filter *FilterParam `form:"filter,omitempty" json:"filter,omitempty"`
}

// GetRolesItemEntitlementsParams defines parameters for GetRolesItemEntitlements.
type GetRolesItemEntitlementsParams struct {
	// Size The number of records to return per response
	Size *PaginationSize `form:"size,omitempty" json:"size,omitempty"`

	// Page The record offset to return results from
	Page *PaginationPage `form:"page,omitempty" json:"page,omitempty"`

	// NextToken The continuation token to retrieve the next set of results
	NextToken *PaginationNextToken `form:"nextToken,omitempty" json:"nextToken,omitempty"`
}

// PostAuthenticationJSONRequestBody defines body for PostAuthentication for application/json ContentType.
type PostAuthenticationJSONRequestBody = IdentityProvider

// PutAuthenticationItemJSONRequestBody defines body for PutAuthenticationItem for application/json ContentType.
type PutAuthenticationItemJSONRequestBody = IdentityProvider

// PostGroupsJSONRequestBody defines body for PostGroups for application/json ContentType.
type PostGroupsJSONRequestBody = Group

// PutGroupsItemJSONRequestBody defines body for PutGroupsItem for application/json ContentType.
type PutGroupsItemJSONRequestBody = Group

// PatchGroupsItemEntitlementsJSONRequestBody defines body for PatchGroupsItemEntitlements for application/json ContentType.
type PatchGroupsItemEntitlementsJSONRequestBody = GroupEntitlementsPatchRequestBody

// PatchGroupsItemIdentitiesJSONRequestBody defines body for PatchGroupsItemIdentities for application/json ContentType.
type PatchGroupsItemIdentitiesJSONRequestBody = GroupIdentitiesPatchRequestBody

// PatchGroupsItemRolesJSONRequestBody defines body for PatchGroupsItemRoles for application/json ContentType.
type PatchGroupsItemRolesJSONRequestBody = GroupRolesPatchRequestBody

// PostIdentitiesJSONRequestBody defines body for PostIdentities for application/json ContentType.
type PostIdentitiesJSONRequestBody = Identity

// PutIdentitiesItemJSONRequestBody defines body for PutIdentitiesItem for application/json ContentType.
type PutIdentitiesItemJSONRequestBody = Identity

// PatchIdentitiesItemEntitlementsJSONRequestBody defines body for PatchIdentitiesItemEntitlements for application/json ContentType.
type PatchIdentitiesItemEntitlementsJSONRequestBody = IdentityEntitlementsPatchRequestBody

// PatchIdentitiesItemGroupsJSONRequestBody defines body for PatchIdentitiesItemGroups for application/json ContentType.
type PatchIdentitiesItemGroupsJSONRequestBody = IdentityGroupsPatchRequestBody

// PatchIdentitiesItemRolesJSONRequestBody defines body for PatchIdentitiesItemRoles for application/json ContentType.
type PatchIdentitiesItemRolesJSONRequestBody = IdentityRolesPatchRequestBody

// PostRolesJSONRequestBody defines body for PostRoles for application/json ContentType.
type PostRolesJSONRequestBody = Role

// PutRolesItemJSONRequestBody defines body for PutRolesItem for application/json ContentType.
type PutRolesItemJSONRequestBody = Role

// PatchRolesItemEntitlementsJSONRequestBody defines body for PatchRolesItemEntitlements for application/json ContentType.
type PatchRolesItemEntitlementsJSONRequestBody = RoleEntitlementsPatchRequestBody

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xdW3PcKBb+KxS7TynF7exkX/ppPUkm49mM7XKc2oeMK0VLdDdjCTSA7PS4+r9vAUJX",
	"kOS+OLZbL3Z3i8vh8PGdAxzQPQxZkjKKqRRweg9TxFGCJeb62y8klphfqN/U1wiLkJNUEkbhFJ4AITmh",
	"CyAZmOuEgGORxVKA2QoGkKhEf2WYqy8UJRhOoUkHAyjCJU6QKlSuUvXElAXX6wBeoAWhSNVyhr/LK3aD",
	"abv2qyUGIaOS0EwnBVKlU7JwLDnBtxjIJQYUf5dAYAnY3ErnEY0WdQ2V7gItsFswjkPGI8Dmc1W1kSnj",
	"tNDPnLPEI0aqCq1KkBBKkiyB0zeBlYZQiReYN8T5TP72iEOzZIa5UYESTFRESk23pYwK7BFJqIJrIqHv",
	"uUjHx0G3gOsA2tI1on5G0SX+K8NCqm+qAzHVH1GaxiTUDZn8KZjucfwdJWmsG/UtJvRGl6D6Sf1fcjyH",
	"U6iV8C3BEhn4qi45DozM6oNkEsVwerwOYIRUoq/XAUywEDqlkgfwXKAAColkJuD07bFKXzb4n6auf0zK",
	"wTIxT8Xk0ipPN7aufFW6be46gGdM/sIyGj2Npp8xCeZanGrD3+6k4WXZ6wB+oSiTS8bJ3/iJNL0mUbX1",
	"b3bS+lrx+ukcZbF8Km3H31McShwBzDnjlfb/e0ewb1WxLorVDTrJ5BJTmTf+grNbEmGuW8NZirkkhixI",
	"5KBgS0sublYjmXCFsq8q83VBSGz2Jw71GHTXLdqVG93dQyJxIvq04WnRuhAAcY5WLRl1HS4p36EUzUhM",
	"rDA5LgScfrWSfb2HmEYpIwpNcEIiVb9Or3pbLlmkUsOPH67g9fp6fR1s075CntUO2rRqK7tsiaO/i9ZU",
	"hMVUWRzTvABenH9W/95/+PTh6oP6fnL17teKEGVZnaIXUpR1uhryQSl61eoWhVaYZUQxSp5HjXbO4lhB",
	"QXVAzbnS6XPJHJJWBdMPAx+kjTz6b4yTnF2a+i0efjP5HXrWiVbfPMOreO7J39JlmbhedNCWZlCrTiVO",
	"OlvWB+K2npxC589cMn3E0sMfBR0qYo/j87mGxDACDTbhFgHX10aiKlfsUYwaJdm6K8rctu4t2MnRsZuR",
	"VN6qj5xl6Qtu2mlhK/aIl7KSeuWXLN5rvbr8epX7rM5UUNT3o1Sb/7R6mai1rftBvbl6DJpv1dWS4nFH",
	"ziUWLOPhfmu0dZS1shi/TAw/Zu8pFO94OqVTOX2iprG8QDJcWmdtWCPdvl67A1la9f1RpFxtjhN2ix3O",
	"fqMBLPX1jbMB+bLNzyxyTFdSlQKLwVjq0FEfqGxVXtWXxqCm+GbX42LS0gLANlotCg7cCnYLuXvlurSw",
	"tWr1kOrQ6oMUF0DO4gEDTafq1mYp1+4V2Wjz1jrcdmXFcNnGiw4VV2k7OawR3loUR0ehKMLRz80VBchu",
	"aiWVQApV1rmaEHom6gkisfPJnHAhz3wT/EXRYc11/MBnPP5khGL3oxgJ+YktCPU+9QqSYp4QIQijHmm4",
	"NnbOR8ajaOoSf5eYUxTrLY78Yy+5GT0WRQZFR3V173M2hd427JRnujW1DePU5wkdzL2wHtJOjaEp1cvd",
	"Dun2otdm63ehUf8iPQpDnEpxwVmSyjNGqyN6xliMEVUFoTBkGZWfCL0hdHFO45U7XRgTBfr3zs4xDz/j",
	"kGP3UnFEBJrF2Ep9SufMXQ2mKl3kfujhOuvrvFMNcZOPd+2U44hwHMov3E3MQjKO9fazcMtUSXCJUaSk",
	"9yRc0fB3FuEqjEmSMi7dC+I8E/JDw2AUxa0HQGJXlnX7XZPaRPlp+W1t0fYy9Hfovdm5uWfJfTVs/twx",
	"nOxYqZppdCcmOHt9h4UbrSniAxb6C9Ede4J5vXZLorPp2yK7lGNTRFfXC+qClNvDQ1YPPunE1X3jIbl+",
	"V2nXlc1jJ3GZveN7V3BKtY15QitCYFtQFt+lgE+2tXUt2K3xwTo4Uxlaaxvqx97az/K66hKYTfk+ctCp",
	"umr43W7nN2igpvWKoVFPisiodqfkoUDtfHlwQH9vqRKcAudk6N2FG05ajbW+9igJ9rRk1ay4b1PRsxM6",
	"nAF5hUvbMvfJ9+ymEU75d2rv/BrayuTZWeU2nK/Gx4Z8r5KR3GUNGZUo1PDLp/Mwzm7w0R0WMV69XrI4",
	"xqv/hIgySkIUH4U6ojCP2PuU3WDwP5PyV50StoJ0rpYYzFkcsztCF0CkODRLCoRRwDIZE4qFDqE8uTgF",
	"VngwZ1z/+MvHE4CihFAiJDeZ5lzHNkVAMqDn2CiU4I7IJUAUnKeYqjyEColoiIFccpYtlgCBlLMoC6VQ",
	"FR2BqyURgAiVB9+y+LYtHBKA43ls4osI1eLcYi7UMxPpePQHhQHU6IBT+M7qqBDipC74hREAvGNJiiQx",
	"MSpKGhjAvGA4hcdHx0dvjs0UEVOUEjiFPx29OTqGio3lUiNjgmqb9XqqaaYqTeUTAWy4CYiJkAKgOAb5",
	"tETpxPjWpikKjbrA0whO21EIsB5a4qGGMsmkETTq21pw5tBRrw/KUUbxrq8b4aD/Oj4eEBU3LCKtc7/O",
	"EaV2/l/VmW+NBK6CC0knlahVneVNf5ZmDODb47f9mYoA0XrQYHcmm1AH2WVJgvhKUQAREoSMzskiUwO3",
	"DswKwNRIQQvjfddRpag8ZcIB33e2XIAAxXe+wo9a0L1goo1dXrcMOwFDe1JZ52DJM7zeIxjd9R8SAB+A",
	"ES/+1kGTUSdpddVhALeKJcviyEa7I821gM1t/QLcUHanzw0oOyJWQuIk0GYun5eC8xuJAvCRsUWMA/CR",
	"yF+zGcAyPPojOz7+KZxxMNGf8B/01aur8/fnr15NwWvwWRutlS42r12Z+yF8Xq6sjMRuiL0v6u7gRtel",
	"RpRx0CykRZamjGuvKGefgTTvGGb3JFqboRVj6ThWcqnnCMpL84zrwvbon9uW4L0uuN6v2oF3A+mw+vbB",
	"uu2w4E6O/IglQEAQuojxcOvdGogdHTZa0Z1gYWhHdblwdQuiT3ipCUs5XSQRbPpGXWfgrgOYZg5QfUkj",
	"JLfhhIvMh6/RQXx50N4ALp0WLGycmRngHN6ROAYzDDKBIzAzvlp1GSNSwyYhFIO7JQmX4MspwPmKk3Us",
	"Z7iYtJdLGglKU0IXR+AkjgG6RSRWKYp6q5mV8cYRWGLlKMeMLsySiZKkzJifRnEyci0uf79emPOkgQe3",
	"vd5KqQyi3GylVNsJxKwFVTq81remu5tLzp7uxqDZ5OK4q/bHDQasg77E4LfP52cgHxnCADJhEY6dyq+u",
	"QT4fZ70/W/Xg9759e2e470Ea+vrwKLUCCDWwnDOeIFkdGLVR0B4YE47uOgeHxN/lJI0R6RsWHN0NHhWX",
	"6K4xMHrwU0pRR07T8xgxUWICCd0nSnU9gCjD+vxzgTgGJpmzQ/NIzpHgNiK4xqmVw5zD1ABmwZojs2Oh",
	"mWPtI+oVRJ3avaZcAHQfU4U8BPlx5weVSg9q1djZ4S28lLQ2eKXKBx+zEGUANC5A1RagOjpgyLqSR+EF",
	"I+579WikXk9/uOj3sReJvGyeNdExEvpLWOUZTubDZvZ6f9vEVNf9YuUw69LAEoke9nme0/fHcFfHSfnq",
	"QSB7LEpFMlw6ruKLIsA44NZuO0VmkzlnSQfzqrK9Y+NhPFzGdOvTbSc0Mj6FBmKh3hNabLKzOZgxEyqG",
	"ogggGtnWFELCAN6iOMO1eMGvrWDN9lU31Z9eq59eo8aVNHmS1Wv1rfq0WsDK5s1PCeloSTWGNxBg1inA",
	"rFOAWSFAHqi5vl6vB9/b1X/aed2+te8AfeANh1SvcSO1A6leJ1oRTJlUDQ80yJ8+rV4DNtqzSceVNwe7",
	"nOjF1VMyYhUhH2q5amPg6dqt8k6G4qM2PTXb4ko0g9vyf8d1DMNnRYdrDbqx2WsCisPrneyvUw0nfnN2",
	"YeT8CufXb9w5WLp3AelpTVeUhMVIKqZWvTxvIf90Kb6gcnuIWP/TU4j7GoXXn882oHT3weKRzYf49n78",
	"Obl8oAtvd23L5E4Kf44e+9PaXx19+y7AWQhXUOvfbz3RBFfEj7s3W7fwrx8WG/mjYjIPDz6Ofnfipk5/",
	"Dzgl4EeU2X8tMTXuwTYOAfR2ypC9WL/+awT6WAH9B74TO6BLf0DUfgfrZy6MjMT/koLwN+H+B2zVtrbP",
	"ijNqvm3aOuLGrdqO8+7jbu1qMM4ek2+HLH9UBC5XEzvJWN910jU4xr3al7FXO+g6znG7drXJmBps4nqi",
	"5zXpmDR1uiECIJDg/M17AyzccwuyfyTbdughnA8A2FMzbTVx7Tj0HTNxmLUNo/of06DlFwmb/+3t2/rj",
	"rTZue24PHs3A6qH4G2wCujds9QA1m2xICLKgOLLXyTxg1WXcv3WR/4Fv4Q6E1rPgfd2MQbQ/7uv6LnUe",
	"CX8Dwi+A5+V7Xr3vuTMox26xlVczFHndJ6GLkp/T3q7rJeBGqVfmdZ4lfh0nylcpLg4rryovh9cvhXdQ",
	"yj7tSPslZgd/iNwDXTs4yrGQj41BwWq2cA/J29efjREOG8J4jGRzoqwArQbXwGPkKrE7sGEzx2Pg+9L1",
	"vdC739ba5lLqQz5bblHQhFDBeb3BDSZwIQ8f88U1aEiNIQ1a/y2FtYdvfyCDW9mWI/cdv1CO4oONXfB3",
	"3uOfIfdQedYAw/Nh8xF73vPjwwh7WESCdShqm8U6Qr6fXsZABI+DOgYhrAZB65GIcxcnxT38qpeiPONh",
	"jD14GbEHve/wGdcfVxuPJ5cNE3doscD8yA6NAdfX8sq1qtVLVM9TTE8uTvV7c4L86trQXjir77qdMw4i",
	"PMsWC0IXAZhxdif0J5UtIiJkt5ivAv2ag9b4/2wk/U3kb+/YwqwMuuDRe5VstZ0ACYDyWzJJXcv6fXM5",
	"9jG/tTRbryiNUYiXLI4whwHMeAyncCllKqaTSeXZUciSye0bbVDz8lsUS8se0ncIq4czszhdvcdWcVn+",
	"fqWS6GsX3SrScOw1ojhuXpmcXxBcLu6VJTbuSm6XeQJCFsc4NG9+qh4fqdqf4rf+AhZ22z7PnH/vz8jz",
	"9Zc8n/nanw3XLVCVKO2vA+qurNTb+ouf1tfr/wcAAP//g/SxrimVAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
