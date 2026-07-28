---
title: Get exemption rules for an account
page_id: operation-get-accounts-account-id-cloudforce-one-rules-exemptions-92cd85a2
path: operations/rules
description: Get all exemption rule patterns for the account, grouped by type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/exemptions
operation_ids:
    - cloudforce-one-get-exemptions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get exemption rules for an account

`GET /accounts/{account_id}/cloudforce-one/rules/exemptions`

Operation ID: `cloudforce-one-get-exemptions`

Get all exemption rule patterns for the account, grouped by type.

## Definition

```yaml
{"operationId": "cloudforce-one-get-exemptions", "summary": "Get exemption rules for an account", "description": "Get all exemption rule patterns for the account, grouped by type.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "responses": {"200": {"description": "Exemption rules grouped by type.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_AccountExemptions"}}}}, "400": {"description": "Validation error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
