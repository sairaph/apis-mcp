---
title: Get a PAC file
page_id: operation-get-accounts-account-id-gateway-pacfiles-pacfile-id-2dd1bf83
path: operations/zero-trust-gateway-pac-files
description: Get a single Zero Trust Gateway PAC file.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/pacfiles/{pacfile_id}
operation_ids:
    - zero-trust-gateway-pacfiles-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a PAC file

`GET /accounts/{account_id}/gateway/pacfiles/{pacfile_id}`

Operation ID: `zero-trust-gateway-pacfiles-details`

Get a single Zero Trust Gateway PAC file.

## Definition

```yaml
{"operationId": "zero-trust-gateway-pacfiles-details", "summary": "Get a PAC file", "description": "Get a single Zero Trust Gateway PAC file.", "parameters": [{"name": "pacfile_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Returns a PAC file response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-7"}}}}, "4XX": {"description": "Returns a PAC file response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-7"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway PAC files"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.pacfiles", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
