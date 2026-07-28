---
title: List account configuration
page_id: operation-get-accounts-account-id-mnm-config-e5fcd468
path: operations/magic-network-monitoring-configuration
description: Lists default sampling, router IPs and warp devices for account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/mnm/config
operation_ids:
    - magic-network-monitoring-configuration-list-account-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account configuration

`GET /accounts/{account_id}/mnm/config`

Operation ID: `magic-network-monitoring-configuration-list-account-configuration`

Lists default sampling, router IPs and warp devices for account.

## Definition

```yaml
{"operationId": "magic-network-monitoring-configuration-list-account-configuration", "summary": "List account configuration", "description": "Lists default sampling, router IPs and warp devices for account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "responses": {"200": {"description": "List account configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}}}}, "4XX": {"description": "List account configuration response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Configuration"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write", "Magic Network Monitoring Config Read"]}
```
