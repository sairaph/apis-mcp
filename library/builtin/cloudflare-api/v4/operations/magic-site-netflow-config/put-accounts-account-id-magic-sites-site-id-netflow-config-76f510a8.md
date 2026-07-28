---
title: Update NetFlow Configuration
page_id: operation-put-accounts-account-id-magic-sites-site-id-netflow-config-ece54d9a
path: operations/magic-site-netflow-config
description: Updates NetFlow configuration for a site (partial update).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/netflow_config
operation_ids:
    - magic-site-netflow-config-update-netflow-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update NetFlow Configuration

`PUT /accounts/{account_id}/magic/sites/{site_id}/netflow_config`

Operation ID: `magic-site-netflow-config-update-netflow-config`

Updates NetFlow configuration for a site (partial update).

## Definition

```yaml
{"operationId": "magic-site-netflow-config-update-netflow-config", "summary": "Update NetFlow Configuration", "description": "Updates NetFlow configuration for a site (partial update).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_netflow_config_request"}}}}, "responses": {"200": {"description": "Update NetFlow Configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_netflow_config_single_response"}}}}, "4XX": {"description": "Update NetFlow Configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site NetFlow Config"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.netflow-config", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
