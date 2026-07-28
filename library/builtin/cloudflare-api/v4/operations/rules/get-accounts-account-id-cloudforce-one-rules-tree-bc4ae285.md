---
title: Get folder tree structure
page_id: operation-get-accounts-account-id-cloudforce-one-rules-tree-e311bc0a
path: operations/rules
description: Get the folder tree structure for rules navigation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/tree
operation_ids:
    - cloudforce-one-get-rule-tree
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get folder tree structure

`GET /accounts/{account_id}/cloudforce-one/rules/tree`

Operation ID: `cloudforce-one-get-rule-tree`

Get the folder tree structure for rules navigation.

## Definition

```yaml
{"operationId": "cloudforce-one-get-rule-tree", "summary": "Get folder tree structure", "description": "Get the folder tree structure for rules navigation.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "responses": {"200": {"description": "Folder tree structure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_TreeResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
