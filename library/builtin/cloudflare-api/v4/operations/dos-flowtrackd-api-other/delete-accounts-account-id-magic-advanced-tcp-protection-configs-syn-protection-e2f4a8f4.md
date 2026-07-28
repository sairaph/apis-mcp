---
title: Delete SYN Protection rule.
page_id: operation-delete-accounts-account-id-magic-advanced-tcp-protection-configs-syn-pro-207eb458
path: operations/dos-flowtrackd-api-other
description: Delete a SYN Protection rule specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_tcp_protection/configs/syn_protection/rules/{rule_id}
operation_ids:
    - deleteSynProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete SYN Protection rule.

`DELETE /accounts/{account_id}/magic/advanced_tcp_protection/configs/syn_protection/rules/{rule_id}`

Operation ID: `deleteSynProtectionRule`

Delete a SYN Protection rule specified by the given UUID.

## Definition

```yaml
{"operationId": "deleteSynProtectionRule", "summary": "Delete SYN Protection rule.", "description": "Delete a SYN Protection rule specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "rule_id", "in": "path", "description": "The UUID of the SYN Protection rule to delete.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Delete SYN Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}, "4XX": {"description": "Delete SYN Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
