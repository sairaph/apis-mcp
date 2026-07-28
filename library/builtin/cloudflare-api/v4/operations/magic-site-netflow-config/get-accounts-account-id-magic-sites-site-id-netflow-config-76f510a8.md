---
title: NetFlow Configuration Details
page_id: operation-get-accounts-account-id-magic-sites-site-id-netflow-config-4f1742c0
path: operations/magic-site-netflow-config
description: Get NetFlow configuration for a site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/sites/{site_id}/netflow_config
operation_ids:
    - magic-site-netflow-config-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# NetFlow Configuration Details

`GET /accounts/{account_id}/magic/sites/{site_id}/netflow_config`

Operation ID: `magic-site-netflow-config-details`

Get NetFlow configuration for a site.

## Definition

```yaml
{"operationId": "magic-site-netflow-config-details", "summary": "NetFlow Configuration Details", "description": "Get NetFlow configuration for a site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Get NetFlow Configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_netflow_config_single_response"}}}}, "4XX": {"description": "Get NetFlow Configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Site NetFlow Config"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.sites.netflow-config", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
