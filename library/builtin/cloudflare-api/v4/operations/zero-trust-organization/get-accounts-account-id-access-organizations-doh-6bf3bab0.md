---
title: Get your Zero Trust organization DoH settings
page_id: operation-get-accounts-account-id-access-organizations-doh-004a7646
path: operations/zero-trust-organization
description: Returns the DoH settings for your Zero Trust organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/organizations/doh
operation_ids:
    - zero-trust-organization-get-your-zero-trust-organization-doh-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get your Zero Trust organization DoH settings

`GET /accounts/{account_id}/access/organizations/doh`

Operation ID: `zero-trust-organization-get-your-zero-trust-organization-doh-settings`

Returns the DoH settings for your Zero Trust organization.

## Definition

```yaml
{"operationId": "zero-trust-organization-get-your-zero-trust-organization-doh-settings", "summary": "Get your Zero Trust organization DoH settings", "description": "Returns the DoH settings for your Zero Trust organization.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get your Zero Trust organization DoH settings response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/access_single_response-2"}, {"properties": {"result": {"type": "object", "properties": {"doh_jwt_duration": {"description": "The duration the DoH JWT is valid for. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.  Note that the maximum duration for this setting is the same as the key rotation period on the account.", "type": "string", "example": "800h"}}}}}]}}}}, "4XX": {"description": "Get your Zero Trust organization DoH settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Revoke", "Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.organizations.doh", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
