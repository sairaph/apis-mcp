---
title: Create new target
page_id: operation-post-accounts-account-id-infrastructure-targets-c54228cd
path: operations/infrastructure-access-targets
description: Create new target
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/infrastructure/targets
operation_ids:
    - infra-targets-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new target

`POST /accounts/{account_id}/infrastructure/targets`

Operation ID: `infra-targets-post`

## Definition

```yaml
{"operationId": "infra-targets-post", "summary": "Create new target", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"hostname": {"description": "A non-unique field that refers to a target. Case insensitive, maximum\nlength of 255 characters, supports the use of special characters dash\nand period, does not support spaces, and must start and end with an\nalphanumeric character.", "type": "string", "example": "infra-access-target", "x-auditable": true}, "ip": {"$ref": "#/components/schemas/infra_IPInfo"}}, "required": ["hostname", "ip"]}}}}, "responses": {"200": {"description": "Successfully created the target", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_Target"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to create the target", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Infrastructure Access Targets"]}
```
