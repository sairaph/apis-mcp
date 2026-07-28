---
title: Update exemption rule patterns
page_id: operation-put-accounts-account-id-cloudforce-one-rules-exemptions-b702db02
path: operations/rules
description: Replace existing exemption patterns with new values. Each key maps to an array of {old_pattern, new_pattern} entries. Missing keys leave that type untouched. Fails if any old pattern is not found or any new pattern already exists.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/exemptions
operation_ids:
    - cloudforce-one-update-account-exemptions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update exemption rule patterns

`PUT /accounts/{account_id}/cloudforce-one/rules/exemptions`

Operation ID: `cloudforce-one-update-account-exemptions`

Replace existing exemption patterns with new values. Each key maps to an array of {old_pattern, new_pattern} entries. Missing keys leave that type untouched. Fails if any old pattern is not found or any new pattern already exists.

## Definition

```yaml
{"operationId": "cloudforce-one-update-account-exemptions", "summary": "Update exemption rule patterns", "description": "Replace existing exemption patterns with new values. Each key maps to an array of {old_pattern, new_pattern} entries. Missing keys leave that type untouched. Fails if any old pattern is not found or any new pattern already exists.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_UpdateAccountExemptionsBody"}}}}, "responses": {"200": {"description": "Full exemption state after the updates.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_AccountExemptions"}}}}, "400": {"description": "Validation error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "404": {"description": "Pattern not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "409": {"description": "New pattern already exists.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
