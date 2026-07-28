---
title: Get allowlist prefix.
page_id: operation-get-accounts-account-id-magic-advanced-tcp-protection-configs-allowlist-3a6b9313
path: operations/dos-flowtrackd-api-other
description: Get an allowlist prefix specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/allowlist/{prefix_id}
operation_ids:
    - getAllowlistPrefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get allowlist prefix.

`GET /accounts/{account_id}/magic/advanced_tcp_protection/configs/allowlist/{prefix_id}`

Operation ID: `getAllowlistPrefix`

Get an allowlist prefix specified by the given UUID.

## Definition

```yaml
{"operationId": "getAllowlistPrefix", "summary": "Get allowlist prefix.", "description": "Get an allowlist prefix specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "prefix_id", "in": "path", "description": "The UUID of the allowlist prefix.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Get allowlist prefix response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_infra-prefix-response"}}}}, "4XX": {"description": "Get allowlist prefix failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write", "DDoS Protection Read"]}
```
