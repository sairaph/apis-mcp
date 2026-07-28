---
title: Update an IP Access rule
page_id: operation-patch-accounts-account-id-firewall-access-rules-rules-rule-id-4e2bf3bf
path: operations/ip-access-rules-for-an-account
description: |-
    Updates an IP Access rule defined at the account level.

    Note: This operation will affect all zones in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/firewall/access_rules/rules/{rule_id}
operation_ids:
    - ip-access-rules-for-an-account-update-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an IP Access rule

`PATCH /accounts/{account_id}/firewall/access_rules/rules/{rule_id}`

Operation ID: `ip-access-rules-for-an-account-update-an-ip-access-rule`

Updates an IP Access rule defined at the account level.

Note: This operation will affect all zones in the account.

## Definition

```yaml
{"operationId": "ip-access-rules-for-an-account-update-an-ip-access-rule", "summary": "Update an IP Access rule", "description": "Updates an IP Access rule defined at the account level.\n\nNote: This operation will affect all zones in the account.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rule_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_schemas-rule"}}}}, "responses": {"200": {"description": "Update an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_response_single"}}}}, "4XX": {"description": "Update an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for an account"], "x-api-token-group": ["Account Firewall Access Rules Write"], "x-cfPermissionsRequired": {"enum": ["#waf:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.access-rules", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
