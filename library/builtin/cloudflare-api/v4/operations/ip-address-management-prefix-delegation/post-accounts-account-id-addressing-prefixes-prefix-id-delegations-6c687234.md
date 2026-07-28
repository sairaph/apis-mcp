---
title: Create Prefix Delegation
page_id: operation-post-accounts-account-id-addressing-prefixes-prefix-id-delegations-5c5e2c41
path: operations/ip-address-management-prefix-delegation
description: Create a new account delegation for a given IP prefix.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations
operation_ids:
    - ip-address-management-prefix-delegation-create-prefix-delegation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Prefix Delegation

`POST /accounts/{account_id}/addressing/prefixes/{prefix_id}/delegations`

Operation ID: `ip-address-management-prefix-delegation-create-prefix-delegation`

Create a new account delegation for a given IP prefix.

## Definition

```yaml
{"operationId": "ip-address-management-prefix-delegation-create-prefix-delegation", "summary": "Create Prefix Delegation", "description": "Create a new account delegation for a given IP prefix.", "parameters": [{"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"cidr": {"$ref": "#/components/schemas/addressing_cidr"}, "delegated_account_id": {"$ref": "#/components/schemas/addressing_delegated_account_identifier"}}, "required": ["cidr", "delegated_account_id"]}}}}, "responses": {"200": {"description": "Create Prefix Delegation response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_schemas-single_response"}}}}, "4XX": {"description": "Create Prefix Delegation response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_schemas-single_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefix Delegation"], "x-api-token-group": ["IP Prefixes: Write"]}
```
