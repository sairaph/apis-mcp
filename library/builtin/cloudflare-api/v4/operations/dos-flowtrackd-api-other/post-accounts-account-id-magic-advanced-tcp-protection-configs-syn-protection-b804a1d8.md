---
title: Create a SYN Protection filter.
page_id: operation-post-accounts-account-id-magic-advanced-tcp-protection-configs-syn-prote-e833ebba
path: operations/dos-flowtrackd-api-other
description: Create a SYN Protection filter for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/syn_protection/filters
operation_ids:
    - createSynProtectionFilter
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a SYN Protection filter.

`POST /accounts/{account_id}/magic/advanced_tcp_protection/configs/syn_protection/filters`

Operation ID: `createSynProtectionFilter`

Create a SYN Protection filter for an account.

## Definition

```yaml
{"operationId": "createSynProtectionFilter", "summary": "Create a SYN Protection filter.", "description": "Create a SYN Protection filter for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "requestBody": {"description": "The new filter to create.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_NewExpressionFilter"}}}}, "responses": {"200": {"description": "Create SYN Protection filter response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_expression-filter-response"}}}}, "4XX": {"description": "Create SYN Protection filter failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
