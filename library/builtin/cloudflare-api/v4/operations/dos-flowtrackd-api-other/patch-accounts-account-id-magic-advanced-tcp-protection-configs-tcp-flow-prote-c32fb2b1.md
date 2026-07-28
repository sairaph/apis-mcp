---
title: Update TCP Flow Protection filter.
page_id: operation-patch-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flow-61a7f634
path: operations/dos-flowtrackd-api-other
description: Update a TCP Flow Protection filter specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters/{filter_id}
operation_ids:
    - updateTcpFlowProtectionFilter
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update TCP Flow Protection filter.

`PATCH /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters/{filter_id}`

Operation ID: `updateTcpFlowProtectionFilter`

Update a TCP Flow Protection filter specified by the given UUID.

## Definition

```yaml
{"operationId": "updateTcpFlowProtectionFilter", "summary": "Update TCP Flow Protection filter.", "description": "Update a TCP Flow Protection filter specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "filter_id", "in": "path", "description": "The UUID of the filter to update.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "requestBody": {"description": "The updates to apply to the filter.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_ExpressionFilterUpdate"}}}}, "responses": {"200": {"description": "Update TCP Flow Protection filter response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_expression-filter-response"}}}}, "4XX": {"description": "Update TCP Flow Protection filter failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
