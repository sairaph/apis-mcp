---
title: Create multiple prefixes.
page_id: operation-post-accounts-account-id-magic-advanced-tcp-protection-configs-prefixes-fad8521e
path: operations/dos-flowtrackd-api-other
description: Create multiple prefixes for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes/bulk
operation_ids:
    - bulkCreatePrefixes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create multiple prefixes.

`POST /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes/bulk`

Operation ID: `bulkCreatePrefixes`

Create multiple prefixes for an account.

## Definition

```yaml
{"operationId": "bulkCreatePrefixes", "summary": "Create multiple prefixes.", "description": "Create multiple prefixes for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "requestBody": {"description": "The list of new prefixes to create.", "required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/dos_NewPrefix"}}}}}, "responses": {"200": {"description": "Create multiple prefixes response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_prefix-list-response"}}}}, "4XX": {"description": "Create multiple prefixes failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
