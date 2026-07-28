---
title: Delete allowlist prefix.
page_id: operation-delete-accounts-account-id-magic-advanced-tcp-protection-configs-allowli-2bb55db1
path: operations/dos-flowtrackd-api-other
description: Delete the allowlist prefix for an account given a UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/allowlist/{prefix_id}
operation_ids:
    - deleteAllowlistPrefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete allowlist prefix.

`DELETE /accounts/{account_id}/magic/advanced_tcp_protection/configs/allowlist/{prefix_id}`

Operation ID: `deleteAllowlistPrefix`

Delete the allowlist prefix for an account given a UUID.

## Definition

```yaml
{"operationId": "deleteAllowlistPrefix", "summary": "Delete allowlist prefix.", "description": "Delete the allowlist prefix for an account given a UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "prefix_id", "in": "path", "description": "The UUID of the allowlist prefix to delete.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Delete allowlist prefix response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}, "4XX": {"description": "Delete allowlist prefix failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
