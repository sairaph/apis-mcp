---
title: Add patterns to exemption rules
page_id: operation-post-accounts-account-id-cloudforce-one-rules-exemptions-bd839a28
path: operations/rules
description: Add regex patterns to per-account exemption rules (union semantics). Missing keys leave that type untouched; duplicates are silently deduped.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/exemptions
operation_ids:
    - cloudforce-one-add-account-exemptions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add patterns to exemption rules

`POST /accounts/{account_id}/cloudforce-one/rules/exemptions`

Operation ID: `cloudforce-one-add-account-exemptions`

Add regex patterns to per-account exemption rules (union semantics). Missing keys leave that type untouched; duplicates are silently deduped.

## Definition

```yaml
{"operationId": "cloudforce-one-add-account-exemptions", "summary": "Add patterns to exemption rules", "description": "Add regex patterns to per-account exemption rules (union semantics). Missing keys leave that type untouched; duplicates are silently deduped.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_PartialAccountExemptions"}}}}, "responses": {"200": {"description": "Full exemption state after the union.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_AccountExemptions"}}}}, "400": {"description": "Validation error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
