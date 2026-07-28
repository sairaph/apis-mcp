---
title: Remove patterns from exemption rules
page_id: operation-delete-accounts-account-id-cloudforce-one-rules-exemptions-7d489124
path: operations/rules
description: Remove regex patterns from per-account exemption rules. Missing keys leave that type untouched; non-existent patterns are silently skipped.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/exemptions
operation_ids:
    - cloudforce-one-remove-account-exemptions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove patterns from exemption rules

`DELETE /accounts/{account_id}/cloudforce-one/rules/exemptions`

Operation ID: `cloudforce-one-remove-account-exemptions`

Remove regex patterns from per-account exemption rules. Missing keys leave that type untouched; non-existent patterns are silently skipped.

## Definition

```yaml
{"operationId": "cloudforce-one-remove-account-exemptions", "summary": "Remove patterns from exemption rules", "description": "Remove regex patterns from per-account exemption rules. Missing keys leave that type untouched; non-existent patterns are silently skipped.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_PartialAccountExemptions"}}}}, "responses": {"200": {"description": "Full exemption state after the subtraction.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_AccountExemptions"}}}}, "400": {"description": "Validation error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
