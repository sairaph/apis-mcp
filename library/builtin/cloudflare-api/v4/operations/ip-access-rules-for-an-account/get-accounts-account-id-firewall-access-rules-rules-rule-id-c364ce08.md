---
title: Get an IP Access rule
page_id: operation-get-accounts-account-id-firewall-access-rules-rules-rule-id-ad011f91
path: operations/ip-access-rules-for-an-account
description: Fetches the details of an IP Access rule defined at the account level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/firewall/access_rules/rules/{rule_id}
operation_ids:
    - ip-access-rules-for-an-account-get-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an IP Access rule

`GET /accounts/{account_id}/firewall/access_rules/rules/{rule_id}`

Operation ID: `ip-access-rules-for-an-account-get-an-ip-access-rule`

Fetches the details of an IP Access rule defined at the account level.

## Definition

```yaml
{"operationId": "ip-access-rules-for-an-account-get-an-ip-access-rule", "summary": "Get an IP Access rule", "description": "Fetches the details of an IP Access rule defined at the account level.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rule_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_account_identifier"}}], "responses": {"200": {"description": "Get an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_response_single"}}}}, "4XX": {"description": "Get an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for an account"], "x-api-token-group": ["Account Firewall Access Rules Write", "Account Firewall Access Rules Read"], "x-cfPermissionsRequired": {"enum": ["#waf:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.access-rules", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
