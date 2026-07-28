---
title: Create NetFlow Configuration
page_id: operation-post-accounts-account-id-magic-sites-site-id-netflow-config-083fdb56
path: operations/magic-site-netflow-config
description: Creates a NetFlow configuration for a site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/netflow_config
operation_ids:
    - magic-site-netflow-config-create-netflow-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create NetFlow Configuration

`POST /accounts/{account_id}/magic/sites/{site_id}/netflow_config`

Operation ID: `magic-site-netflow-config-create-netflow-config`

Creates a NetFlow configuration for a site.

## Definition

```yaml
{"operationId": "magic-site-netflow-config-create-netflow-config", "summary": "Create NetFlow Configuration", "description": "Creates a NetFlow configuration for a site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_netflow_config_request"}}}}, "responses": {"201": {"description": "Create NetFlow Configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_netflow_config_single_response"}}}}, "4XX": {"description": "Create NetFlow Configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site NetFlow Config"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.netflow-config", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
