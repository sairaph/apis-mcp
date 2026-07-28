---
title: List Prefix Delegations
page_id: operation-get-accounts-account-id-addressing-prefixes-prefix-id-delegations-786917af
path: operations/ip-address-management-prefix-delegation
description: List all delegations for a given account IP prefix.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations
operation_ids:
    - ip-address-management-prefix-delegation-list-prefix-delegations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Prefix Delegations

`GET /accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations`

Operation ID: `ip-address-management-prefix-delegation-list-prefix-delegations`

List all delegations for a given account IP prefix.

## Definition

```yaml
{"operationId": "ip-address-management-prefix-delegation-list-prefix-delegations", "summary": "List Prefix Delegations", "description": "List all delegations for a given account IP prefix.", "parameters": [{"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"200": {"description": "List Prefix Delegations response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_schemas-response_collection"}}}}, "4XX": {"description": "List Prefix Delegations response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_schemas-response_collection"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefix Delegation"]}
```
