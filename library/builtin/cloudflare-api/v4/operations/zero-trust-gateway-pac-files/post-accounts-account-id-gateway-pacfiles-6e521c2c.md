---
title: Create a PAC file
page_id: operation-post-accounts-account-id-gateway-pacfiles-c8213f28
path: operations/zero-trust-gateway-pac-files
description: Create a new Zero Trust Gateway PAC file.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/pacfiles
operation_ids:
    - zero-trust-gateway-pacfiles-create-pacfile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a PAC file

`POST /accounts/{account_id}/gateway/pacfiles`

Operation ID: `zero-trust-gateway-pacfiles-create-pacfile`

Create a new Zero Trust Gateway PAC file.

## Definition

```yaml
{"operationId": "zero-trust-gateway-pacfiles-create-pacfile", "summary": "Create a PAC file", "description": "Create a new Zero Trust Gateway PAC file.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"contents": {"$ref": "#/components/schemas/zero-trust-gateway_contents"}, "description": {"$ref": "#/components/schemas/zero-trust-gateway_description-4"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-7"}, "slug": {"description": "URL-friendly version of the PAC file name. If not provided, it will be auto-generated", "type": "string", "example": "pac_devops", "x-auditable": true}}, "required": ["name", "contents"]}}}}, "responses": {"200": {"description": "Returns a created PAC file response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-7"}}}}, "4XX": {"description": "Returns a created PAC file response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-7"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway PAC files"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.pacfiles", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
