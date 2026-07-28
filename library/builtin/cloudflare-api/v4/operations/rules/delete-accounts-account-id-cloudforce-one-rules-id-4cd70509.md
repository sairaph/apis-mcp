---
title: Delete a rule
page_id: operation-delete-accounts-account-id-cloudforce-one-rules-id-30e90d1a
path: operations/rules
description: Delete an existing rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/{id}
operation_ids:
    - cloudforce-one-delete-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a rule

`DELETE /accounts/{account_id}/cloudforce-one/rules/{id}`

Operation ID: `cloudforce-one-delete-rule`

Delete an existing rule.

## Definition

```yaml
{"operationId": "cloudforce-one-delete-rule", "summary": "Delete a rule", "description": "Delete an existing rule.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}, {"name": "id", "in": "path", "description": "The unique identifier for the rule.", "required": true, "schema": {"description": "The unique identifier for the rule.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_DeleteRuleBody"}}}}, "responses": {"200": {"description": "Rule deleted (customer accounts).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_SuccessResponse"}}}}, "202": {"description": "Deletion pending approval (internal accounts).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ApprovalPendingResponse"}}}}, "400": {"description": "Validation error (e.g. commit_message exceeds max length).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "404": {"description": "Rule not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
