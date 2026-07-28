---
title: Update an entire account configuration
page_id: operation-put-accounts-account-id-mnm-config-6ca4880a
path: operations/magic-network-monitoring-configuration
description: Update an existing network monitoring configuration, requires the entire configuration to be updated at once.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/mnm/config
operation_ids:
    - magic-network-monitoring-configuration-update-an-entire-account-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an entire account configuration

`PUT /accounts/{account_id}/mnm/config`

Operation ID: `magic-network-monitoring-configuration-update-an-entire-account-configuration`

Update an existing network monitoring configuration, requires the entire configuration to be updated at once.

## Definition

```yaml
{"operationId": "magic-network-monitoring-configuration-update-an-entire-account-configuration", "summary": "Update an entire account configuration", "description": "Update an existing network monitoring configuration, requires the entire configuration to be updated at once.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"default_sampling": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_default_sampling"}, "name": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_name"}, "router_ips": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_router_ips"}, "warp_devices": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_warp_devices"}}, "required": ["name", "default_sampling"]}}}}, "responses": {"200": {"description": "Update an entire account configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}}}}, "4XX": {"description": "Update an entire account configuration response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Configuration"], "x-api-token-group": ["Magic Network Monitoring Admin", "Magic Network Monitoring Config Write"]}
```
