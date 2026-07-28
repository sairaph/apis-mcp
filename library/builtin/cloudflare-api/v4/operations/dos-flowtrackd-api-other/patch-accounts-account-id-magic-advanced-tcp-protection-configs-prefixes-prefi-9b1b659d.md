---
title: Update prefix.
page_id: operation-patch-accounts-account-id-magic-advanced-tcp-protection-configs-prefixes-1c56e2ae
path: operations/dos-flowtrackd-api-other
description: Update a prefix specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes/{prefix_id}
operation_ids:
    - updatePrefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update prefix.

`PATCH /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes/{prefix_id}`

Operation ID: `updatePrefix`

Update a prefix specified by the given UUID.

## Definition

```yaml
{"operationId": "updatePrefix", "summary": "Update prefix.", "description": "Update a prefix specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "prefix_id", "in": "path", "description": "The UUID of the prefix to update.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "requestBody": {"description": "The updates to apply to the prefix.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_PrefixUpdate"}}}}, "responses": {"200": {"description": "Update prefix response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_prefix-response"}}}}, "4XX": {"description": "Update prefix failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
