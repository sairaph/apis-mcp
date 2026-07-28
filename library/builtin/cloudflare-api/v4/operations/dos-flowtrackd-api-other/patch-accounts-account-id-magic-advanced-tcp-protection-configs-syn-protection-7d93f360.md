---
title: Update SYN Protection filter.
page_id: operation-patch-accounts-account-id-magic-advanced-tcp-protection-configs-syn-prot-5f47a539
path: operations/dos-flowtrackd-api-other
description: Update a SYN Protection filter specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/syn_protection/filters/{filter_id}
operation_ids:
    - updateSynProtectionFilter
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update SYN Protection filter.

`PATCH /accounts/{account_id}/magic/advanced_tcp_protection/configs/syn_protection/filters/{filter_id}`

Operation ID: `updateSynProtectionFilter`

Update a SYN Protection filter specified by the given UUID.

## Definition

```yaml
{"operationId": "updateSynProtectionFilter", "summary": "Update SYN Protection filter.", "description": "Update a SYN Protection filter specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "filter_id", "in": "path", "description": "The UUID of the filter to update.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "requestBody": {"description": "The updates to apply to the filter.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_ExpressionFilterUpdate"}}}}, "responses": {"200": {"description": "Update SYN Protection filter response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_expression-filter-response"}}}}, "4XX": {"description": "Update SYN Protection filter failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
