---
title: Get TCP Flow Protection filter.
page_id: operation-get-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flow-p-b7bcae3c
path: operations/dos-flowtrackd-api-other
description: Get a TCP Flow Protection filter specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters/{filter_id}
operation_ids:
    - getTcpFlowProtectionFilter
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get TCP Flow Protection filter.

`GET /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters/{filter_id}`

Operation ID: `getTcpFlowProtectionFilter`

Get a TCP Flow Protection filter specified by the given UUID.

## Definition

```yaml
{"operationId": "getTcpFlowProtectionFilter", "summary": "Get TCP Flow Protection filter.", "description": "Get a TCP Flow Protection filter specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "filter_id", "in": "path", "description": "The UUID of the filter to retrieve.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Get TCP Flow Protection filter response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_expression-filter-response"}}}}, "4XX": {"description": "Get TCP Flow Protection filter failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write", "DDoS Protection Read"]}
```
