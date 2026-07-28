---
title: Get your Zero Trust organization
page_id: operation-get-zones-zone-id-access-organizations-57fab33a
path: operations/zone-level-zero-trust-organization
description: Returns the configuration for your Zero Trust organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/organizations
operation_ids:
    - zone-level-zero-trust-organization-get-your-zero-trust-organization
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get your Zero Trust organization

`GET /zones/{zone_id}/access/organizations`

Operation ID: `zone-level-zero-trust-organization-get-your-zero-trust-organization`

Returns the configuration for your Zero Trust organization.

## Definition

```yaml
{"operationId": "zone-level-zero-trust-organization-get-your-zero-trust-organization", "summary": "Get your Zero Trust organization", "description": "Returns the configuration for your Zero Trust organization.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-4"}}], "responses": {"200": {"description": "Get your Zero Trust organization response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-15"}}}}, "4XX": {"description": "Get your Zero Trust organization response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Revoke", "Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.zone-organizations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
