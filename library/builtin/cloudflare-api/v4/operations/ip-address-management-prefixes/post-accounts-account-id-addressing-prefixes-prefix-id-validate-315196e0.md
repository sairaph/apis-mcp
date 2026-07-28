---
title: Validate Prefix
page_id: operation-post-accounts-account-id-addressing-prefixes-prefix-id-validate-a0c3c106
path: operations/ip-address-management-prefixes
description: Triggers a new prefix validation. The checks are run asynchronously and include IRR, RPKI, and prefix ownership.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/addressing/prefixes/{prefix_id}/validate
operation_ids:
    - ip-address-management-prefixes-validate-prefix
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate Prefix

`POST /accounts/{account_id}/addressing/prefixes/{prefix_id}/validate`

Operation ID: `ip-address-management-prefixes-validate-prefix`

Triggers a new prefix validation. The checks are run asynchronously and include IRR, RPKI, and prefix ownership.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-validate-prefix", "summary": "Validate Prefix", "description": "Triggers a new prefix validation. The checks are run asynchronously and include IRR, RPKI, and prefix ownership.", "parameters": [{"name": "prefix_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_prefix_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"202": {"description": "Validate Prefix response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_single_response"}}}}, "4XX": {"description": "Validate Prefix response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_single_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": ["Magic Transit Write", "IP Prefixes: Write"]}
```
