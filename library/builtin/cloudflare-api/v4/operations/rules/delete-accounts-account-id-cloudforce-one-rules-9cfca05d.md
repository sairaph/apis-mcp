---
title: Delete all rules
page_id: operation-delete-accounts-account-id-cloudforce-one-rules-04acef27
path: operations/rules
description: Delete all rules in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules
operation_ids:
    - cloudforce-one-delete-all-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete all rules

`DELETE /accounts/{account_id}/cloudforce-one/rules`

Operation ID: `cloudforce-one-delete-all-rules`

Delete all rules in an account.

## Definition

```yaml
{"operationId": "cloudforce-one-delete-all-rules", "summary": "Delete all rules", "description": "Delete all rules in an account.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "responses": {"200": {"description": "All rules deleted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_DeleteAllResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
