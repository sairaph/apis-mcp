---
title: Delete Prefix Delegation
page_id: operation-delete-accounts-account-id-addressing-prefixes-prefix-id-delegations-del-889c9531
path: operations/ip-address-management-prefix-delegation
description: Delete an account delegation for a given IP prefix.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations/{delegation_id}
operation_ids:
    - ip-address-management-prefix-delegation-delete-prefix-delegation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Prefix Delegation

`DELETE /accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations/{delegation_id}`

Operation ID: `ip-address-management-prefix-delegation-delete-prefix-delegation`

Delete an account delegation for a given IP prefix.

## Definition

```yaml
{"operationId": "ip-address-management-prefix-delegation-delete-prefix-delegation", "summary": "Delete Prefix Delegation", "description": "Delete an account delegation for a given IP prefix.", "parameters": [{"name": "delegation_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_delegation_identifier"}}, {"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Prefix Delegation response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_id_response"}}}}, "4XX": {"description": "Delete Prefix Delegation response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_id_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefix Delegation"], "x-api-token-group": ["IP Prefixes: Write"]}
```
