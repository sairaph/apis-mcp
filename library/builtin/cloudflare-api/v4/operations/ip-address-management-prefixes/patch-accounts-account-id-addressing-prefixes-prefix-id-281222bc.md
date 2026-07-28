---
title: Update Prefix Description
page_id: operation-patch-accounts-account-id-addressing-prefixes-prefix-id-ec0064ef
path: operations/ip-address-management-prefixes
description: Modify the description for a prefix owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}
operation_ids:
    - ip-address-management-prefixes-update-prefix-description
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Prefix Description

`PATCH /accounts/{account_id}/addressing/prefixes/{prefix_id}`

Operation ID: `ip-address-management-prefixes-update-prefix-description`

Modify the description for a prefix owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-update-prefix-description", "summary": "Update Prefix Description", "description": "Modify the description for a prefix owned by the account.", "parameters": [{"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"description": {"$ref": "#/components/schemas/addressing_description"}}, "required": ["description"]}}}}, "responses": {"200": {"description": "Update Prefix Description response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_single_response"}}}}, "4XX": {"description": "Update Prefix Description response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_single_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": ["Magic Transit Write", "IP Prefixes: Write"]}
```
