---
title: Delete all TCP Flow Protection filters.
page_id: operation-delete-accounts-account-id-magic-advanced-tcp-protection-configs-tcp-flo-fa958a59
path: operations/dos-flowtrackd-api-other
description: Delete all TCP Flow Protection filters for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters
operation_ids:
    - deleteTcpFlowProtectionFiltersForAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete all TCP Flow Protection filters.

`DELETE /accounts/{account_id}/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters`

Operation ID: `deleteTcpFlowProtectionFiltersForAccount`

Delete all TCP Flow Protection filters for an account.

## Definition

```yaml
{"operationId": "deleteTcpFlowProtectionFiltersForAccount", "summary": "Delete all TCP Flow Protection filters.", "description": "Delete all TCP Flow Protection filters for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}], "responses": {"200": {"description": "Delete all TCP Flow Protection filters response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}, "4XX": {"description": "Delete all TCP Flow Protection filters failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
