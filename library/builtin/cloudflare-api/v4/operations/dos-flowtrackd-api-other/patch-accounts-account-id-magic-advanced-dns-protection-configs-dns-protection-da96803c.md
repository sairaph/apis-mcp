---
title: Update DNS Protection rule.
page_id: operation-patch-accounts-account-id-magic-advanced-dns-protection-configs-dns-prot-1b762c7e
path: operations/dos-flowtrackd-api-other
description: Update a DNS Protection rule specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules/{rule_id}
operation_ids:
    - updateDnsProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update DNS Protection rule.

`PATCH /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules/{rule_id}`

Operation ID: `updateDnsProtectionRule`

Update a DNS Protection rule specified by the given UUID.

## Definition

```yaml
{"operationId": "updateDnsProtectionRule", "summary": "Update DNS Protection rule.", "description": "Update a DNS Protection rule specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "rule_id", "in": "path", "description": "The UUID of the DNS Protection rule to update.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "requestBody": {"description": "The updates to apply to the DNS Protection rule.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_DnsProtectionRuleUpdate"}}}}, "responses": {"200": {"description": "Update DNS Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_dns-protection-rule-response"}}}}, "4XX": {"description": "Update DNS Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write"]}
```
