---
title: Update a rule
page_id: operation-put-accounts-account-id-cloudforce-one-rules-id-67158eff
path: operations/rules
description: Update an existing rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/{id}
operation_ids:
    - cloudforce-one-update-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a rule

`PUT /accounts/{account_id}/cloudforce-one/rules/{id}`

Operation ID: `cloudforce-one-update-rule`

Update an existing rule.

## Definition

```yaml
{"operationId": "cloudforce-one-update-rule", "summary": "Update a rule", "description": "Update an existing rule.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}, {"name": "id", "in": "path", "description": "The unique identifier for the rule.", "required": true, "schema": {"description": "The unique identifier for the rule.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_UpdateRule"}}}}, "responses": {"200": {"description": "Rule updated (customer accounts).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_Rule"}}}}, "202": {"description": "Update pending approval (internal accounts).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ApprovalPendingResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "404": {"description": "Rule not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
