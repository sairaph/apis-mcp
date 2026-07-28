---
title: Prefix Details
page_id: operation-get-accounts-account-id-addressing-prefixes-prefix-id-68c483d9
path: operations/ip-address-management-prefixes
description: List a particular prefix owned by the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}
operation_ids:
    - ip-address-management-prefixes-prefix-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Prefix Details

`GET /accounts/{account_id}/addressing/prefixes/{prefix_id}`

Operation ID: `ip-address-management-prefixes-prefix-details`

List a particular prefix owned by the account.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-prefix-details", "summary": "Prefix Details", "description": "List a particular prefix owned by the account.", "parameters": [{"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"200": {"description": "Prefix Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_single_response"}}}}, "4XX": {"description": "Prefix Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_single_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": ["Magic Transit Read", "Magic Transit Write", "IP Prefixes: Write", "IP Prefixes: Read", "IP Prefixes: BGP On Demand Write", "IP Prefixes: BGP On Demand Read"]}
```
