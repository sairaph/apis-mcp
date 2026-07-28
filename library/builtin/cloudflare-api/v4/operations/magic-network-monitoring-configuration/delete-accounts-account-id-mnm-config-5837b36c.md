---
title: Delete account configuration
page_id: operation-delete-accounts-account-id-mnm-config-9d6a7317
path: operations/magic-network-monitoring-configuration
description: Delete an existing network monitoring configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/mnm/config
operation_ids:
    - magic-network-monitoring-configuration-delete-account-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete account configuration

`DELETE /accounts/{account_id}/mnm/config`

Operation ID: `magic-network-monitoring-configuration-delete-account-configuration`

Delete an existing network monitoring configuration.

## Definition

```yaml
{"operationId": "magic-network-monitoring-configuration-delete-account-configuration", "summary": "Delete account configuration", "description": "Delete an existing network monitoring configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete account configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}}}}, "4XX": {"description": "Delete account configuration response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_config_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring Configuration"], "x-api-token-group": ["Magic Network Monitoring Admin"]}
```
