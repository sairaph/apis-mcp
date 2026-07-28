---
title: Site Details
page_id: operation-get-accounts-account-id-magic-sites-site-id-54122a35
path: operations/magic-sites
description: Get a specific Site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}
operation_ids:
    - magic-sites-site-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Site Details

`GET /accounts/{account_id}/magic/sites/{site_id}`

Operation ID: `magic-sites-site-details`

Get a specific Site.

## Definition

```yaml
{"operationId": "magic-sites-site-details", "summary": "Site Details", "description": "Get a specific Site.", "parameters": [{"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the response body will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "Site Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_site_single_response"}}}}, "4XX": {"description": "Site Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Sites"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
