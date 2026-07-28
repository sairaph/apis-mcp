---
title: Delete all prefixes.
page_id: operation-delete-accounts-account-id-magic-advanced-tcp-protection-configs-prefixe-c8d0c566
path: operations/dos-flowtrackd-api-other
description: Delete all prefixes for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes
operation_ids:
    - deletePrefixesForAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete all prefixes.

`DELETE /accounts/{account_id}/magic/advanced_tcp_protection/configs/prefixes`

Operation ID: `deletePrefixesForAccount`

Delete all prefixes for an account.

## Definition

```yaml
{"operationId": "deletePrefixesForAccount", "summary": "Delete all prefixes.", "description": "Delete all prefixes for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "responses": {"200": {"description": "Delete all prefixes response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}, "4XX": {"description": "Delete all prefixes failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
