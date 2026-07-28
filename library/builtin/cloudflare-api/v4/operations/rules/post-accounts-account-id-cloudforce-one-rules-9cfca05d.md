---
title: Create a rule
page_id: operation-post-accounts-account-id-cloudforce-one-rules-d7de765f
path: operations/rules
description: Create a new detection rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules
operation_ids:
    - cloudforce-one-create-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a rule

`POST /accounts/{account_id}/cloudforce-one/rules`

Operation ID: `cloudforce-one-create-rule`

Create a new detection rule.

## Definition

```yaml
{"operationId": "cloudforce-one-create-rule", "summary": "Create a rule", "description": "Create a new detection rule.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_CreateRule"}}}}, "responses": {"201": {"description": "Rule created (customer accounts).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_Rule"}}}}, "202": {"description": "Rule pending approval (internal accounts).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ApprovalPendingResponse"}}}}, "400": {"description": "Validation error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"], "x-stability": "beta"}
```
