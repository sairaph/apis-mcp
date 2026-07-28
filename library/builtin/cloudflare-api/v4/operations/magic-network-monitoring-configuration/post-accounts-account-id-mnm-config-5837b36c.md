---
title: Create account configuration
page_id: operation-post-accounts-account-id-mnm-config-b02acf84
path: operations/magic-network-monitoring-configuration
description: Create a new network monitoring configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/mnm/config
operation_ids:
    - magic-network-monitoring-configuration-create-account-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create account configuration

`POST /accounts/{account_id}/mnm/config`

Operation ID: `magic-network-monitoring-configuration-create-account-configuration`

Create a new network monitoring configuration.

## Definition

```yaml
{"operationId": "magic-network-monitoring-configuration-create-account-configuration", "summary": "Create account configuration", "description": "Create a new network monitoring configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"default_sampling": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_default_sampling"}, "name": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_name"}, "router_ips": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_router_ips"}, "warp_devices": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_warp_devices"}}, "required": ["name", "default_sampling"]}}}}, "responses": {"200": {"description": "Create account configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}}}}, "4XX": {"description": "Create account configuration response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Configuration"], "x-api-token-group": ["Magic Network Monitoring Admin"]}
```
