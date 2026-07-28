---
title: Get DNS Protection rule.
page_id: operation-get-accounts-account-id-magic-advanced-dns-protection-configs-dns-protec-8bc6fe3e
path: operations/dos-flowtrackd-api-other
description: Get a DNS Protection rule specified by the given UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules/{rule_id}
operation_ids:
    - getDnsProtectionRule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DNS Protection rule.

`GET /accounts/{account_id}/magic/advanced_dns_protection/configs/dns_protection/rules/{rule_id}`

Operation ID: `getDnsProtectionRule`

Get a DNS Protection rule specified by the given UUID.

## Definition

```yaml
{"operationId": "getDnsProtectionRule", "summary": "Get DNS Protection rule.", "description": "Get a DNS Protection rule specified by the given UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "The ID of the account.", "required": true, "schema": {"$ref": "#/components/schemas/dos_identifier"}}, {"name": "rule_id", "in": "path", "description": "The UUID of the DNS Protection rule.", "required": true, "schema": {"$ref": "#/components/schemas/dos_uuid"}}], "responses": {"200": {"description": "Get DNS Protection rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_dns-protection-rule-response"}}}}, "4XX": {"description": "Get DNS Protection rule failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dos_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["dos-flowtrackd-api_other"], "x-api-token-group": ["DDoS Protection Write", "DDoS Protection Read"]}
```
