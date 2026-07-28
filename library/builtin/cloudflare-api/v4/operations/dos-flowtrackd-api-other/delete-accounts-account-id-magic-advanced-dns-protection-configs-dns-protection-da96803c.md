---
title: Delete DNS Protection rule.
page_id: operation-delete-accounts-account-id-magic-advanced-dns-protection-configs-dns-pro-92943a12
path: operations/dos-flowtrackd-api-other
description: Delete a DNS Protection rule specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules/{rule_id}
operation_ids:
    - deleteDnsProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete DNS Protection rule.

`DELETE /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules/{rule_id}`

Operation ID: `deleteDnsProtectionRule`

Delete a DNS Protection rule specified by the given UUID.

## Definition

```yaml
{"operationId": "deleteDnsProtectionRule", "summary": "Delete DNS Protection rule.", "description": "Delete a DNS Protection rule specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "rule_id", "in": "path", "description": "The UUID of the DNS Protection rule to delete.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Delete DNS Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common"}}}}, "4XX": {"description": "Delete DNS Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
