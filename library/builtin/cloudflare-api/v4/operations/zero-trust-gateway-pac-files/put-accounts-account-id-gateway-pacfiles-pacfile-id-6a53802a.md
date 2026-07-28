---
title: Update a Zero Trust Gateway PAC file
page_id: operation-put-accounts-account-id-gateway-pacfiles-pacfile-id-55097220
path: operations/zero-trust-gateway-pac-files
description: Update a configured Zero Trust Gateway PAC file.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/gateway/pacfiles/{pacfile_id}
operation_ids:
    - zero-trust-gateway-pacfiles-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Zero Trust Gateway PAC file

`PUT /accounts/{account_id}/gateway/pacfiles/{pacfile_id}`

Operation ID: `zero-trust-gateway-pacfiles-update`

Update a configured Zero Trust Gateway PAC file.

## Definition

```yaml
{"operationId": "zero-trust-gateway-pacfiles-update", "summary": "Update a Zero Trust Gateway PAC file", "description": "Update a configured Zero Trust Gateway PAC file.", "parameters": [{"name": "pacfile_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"contents": {"$ref": "#/components/schemas/zero-trust-gateway_contents"}, "description": {"$ref": "#/components/schemas/zero-trust-gateway_description-4"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-7"}}, "required": ["name", "description", "contents"]}}}}, "responses": {"200": {"description": "Update a Zero Trust Gateway PAC file response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-7"}}}}, "4XX": {"description": "Update a Zero Trust Gateway PAC file response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-7"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway PAC files"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.pacfiles", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
