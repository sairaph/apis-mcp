---
title: Delete prefix.
page_id: operation-delete-accounts-account-id-magic-advanced-tcp-protection-configs-prefixe-155a86ea
path: operations/dos-flowtrackd-api-other
description: Delete the prefix for an account given a UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes/{prefix_id}
operation_ids:
    - deletePrefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete prefix.

`DELETE /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes/{prefix_id}`

Operation ID: `deletePrefix`

Delete the prefix for an account given a UUID.

## Definition

```yaml
{"operationId": "deletePrefix", "summary": "Delete prefix.", "description": "Delete the prefix for an account given a UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "prefix_id", "in": "path", "description": "The UUID of the prefix to delete.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Delete prefix response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}, "4XX": {"description": "Delete prefix failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
