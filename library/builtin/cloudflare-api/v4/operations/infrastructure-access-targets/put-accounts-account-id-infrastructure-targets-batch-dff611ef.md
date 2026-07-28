---
title: Create new targets
page_id: operation-put-accounts-account-id-infrastructure-targets-batch-41ade5d1
path: operations/infrastructure-access-targets
description: Adds one or more targets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/infrastructure/targets/batch
operation_ids:
    - infra-targets-put-batch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new targets

`PUT /accounts/{account_id}/infrastructure/targets/batch`

Operation ID: `infra-targets-put-batch`

Adds one or more targets.

## Definition

```yaml
{"operationId": "infra-targets-put-batch", "summary": "Create new targets", "description": "Adds one or more targets.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"hostname": {"description": "A non-unique field that refers to a target. Case insensitive, maximum\nlength of 255 characters, supports the use of special characters dash\nand period, does not support spaces, and must start and end with an\nalphanumeric character.", "type": "string", "example": "infra-access-target", "x-auditable": true}, "ip": {"$ref": "#/components/schemas/infra_IPInfo"}}, "required": ["hostname", "ip"], "type": "object"}}}}}, "responses": {"200": {"description": "Successfully created the targets", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_TargetArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to create the targets", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Infrastructure Access Targets"]}
```
