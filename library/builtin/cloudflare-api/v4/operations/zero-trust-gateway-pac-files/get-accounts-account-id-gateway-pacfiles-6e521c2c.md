---
title: List PAC files
page_id: operation-get-accounts-account-id-gateway-pacfiles-ae59c0be
path: operations/zero-trust-gateway-pac-files
description: List all Zero Trust Gateway PAC files for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/pacfiles
operation_ids:
    - zero-trust-gateway-pacfiles-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List PAC files

`GET /accounts/{account_id}/gateway/pacfiles`

Operation ID: `zero-trust-gateway-pacfiles-list`

List all Zero Trust Gateway PAC files for an account.

## Definition

```yaml
{"operationId": "zero-trust-gateway-pacfiles-list", "summary": "List PAC files", "description": "List all Zero Trust Gateway PAC files for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Returns a list of PAC files response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-10"}}}}, "4XX": {"description": "Returns a list of PAC files response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-10"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway PAC files"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.pacfiles", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
