---
title: Delete Prefix
page_id: operation-delete-accounts-account-id-addressing-prefixes-prefix-id-37481806
path: operations/ip-address-management-prefixes
description: Delete an unapproved prefix owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}
operation_ids:
    - ip-address-management-prefixes-delete-prefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Prefix

`DELETE /accounts/{account_id}/addressing/prefixes/{prefix_id}`

Operation ID: `ip-address-management-prefixes-delete-prefix`

Delete an unapproved prefix owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-delete-prefix", "summary": "Delete Prefix", "description": "Delete an unapproved prefix owned by the account.", "parameters": [{"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Prefix response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_api-response-single"}}}}, "4XX": {"description": "Delete Prefix response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_api-response-collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": ["Magic Transit Write"]}
```
