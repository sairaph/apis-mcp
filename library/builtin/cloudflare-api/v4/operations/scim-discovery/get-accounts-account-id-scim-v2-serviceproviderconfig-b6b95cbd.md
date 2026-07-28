---
title: Get SCIM Service Provider Config
page_id: operation-get-accounts-account-id-scim-v2-serviceproviderconfig-0a32c183
path: operations/scim-discovery
description: Returns the SCIM 2.0 Service Provider configuration (RFC 7643 Section 5). IdPs use this endpoint to auto-configure their SCIM integration with Cloudflare, discovering which optional features (patch, bulk, filter, etc.) are supported.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/ServiceProviderConfig
operation_ids:
    - scim-service-provider-config-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get SCIM Service Provider Config

`GET /accounts/{account_id}/scim/v2/ServiceProviderConfig`

Operation ID: `scim-service-provider-config-get`

Returns the SCIM 2.0 Service Provider configuration (RFC 7643 Section 5). IdPs use this endpoint to auto-configure their SCIM integration with Cloudflare, discovering which optional features (patch, bulk, filter, etc.) are supported.

## Definition

```yaml
{"operationId": "scim-service-provider-config-get", "summary": "Get SCIM Service Provider Config", "description": "Returns the SCIM 2.0 Service Provider configuration (RFC 7643 Section 5). IdPs use this endpoint to auto-configure their SCIM integration with Cloudflare, discovering which optional features (patch, bulk, filter, etc.) are supported.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "Get SCIM Service Provider Config response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_service_provider_config"}}}}, "4XX": {"description": "Get SCIM Service Provider Config response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Discovery"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
