---
title: Get target
page_id: operation-get-accounts-account-id-infrastructure-targets-target-id-ec831aea
path: operations/infrastructure-access-targets
description: Get target
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/infrastructure/targets/{target_id}
operation_ids:
    - infra-targets-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get target

`GET /accounts/{account_id}/infrastructure/targets/{target_id}`

Operation ID: `infra-targets-get`

## Definition

```yaml
{"operationId": "infra-targets-get", "summary": "Get target", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}, {"name": "target_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_TargetId"}}], "responses": {"200": {"description": "Successfully retrieved the target", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_Target"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to retrieve the target", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Infrastructure Access Targets"]}
```
