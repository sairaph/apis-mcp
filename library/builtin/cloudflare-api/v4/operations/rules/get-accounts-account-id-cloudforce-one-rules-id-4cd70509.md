---
title: Get a rule
page_id: operation-get-accounts-account-id-cloudforce-one-rules-id-f2289de6
path: operations/rules
description: Get a single rule by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/{id}
operation_ids:
    - cloudforce-one-get-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a rule

`GET /accounts/{account_id}/cloudforce-one/rules/{id}`

Operation ID: `cloudforce-one-get-rule`

Get a single rule by ID.

## Definition

```yaml
{"operationId": "cloudforce-one-get-rule", "summary": "Get a rule", "description": "Get a single rule by ID.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}, {"name": "id", "in": "path", "description": "The unique identifier for the rule.", "required": true, "schema": {"description": "The unique identifier for the rule.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}}], "responses": {"200": {"description": "Rule details.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_Rule"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "404": {"description": "Rule not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
