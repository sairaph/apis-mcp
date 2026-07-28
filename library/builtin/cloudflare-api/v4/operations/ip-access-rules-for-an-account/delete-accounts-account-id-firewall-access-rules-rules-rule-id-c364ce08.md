---
title: Delete an IP Access rule
page_id: operation-delete-accounts-account-id-firewall-access-rules-rules-rule-id-2ae97447
path: operations/ip-access-rules-for-an-account
description: |-
    Deletes an existing IP Access rule defined at the account level.

    Note: This operation will affect all zones in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/firewall/access_rules/rules/{rule_id}
operation_ids:
    - ip-access-rules-for-an-account-delete-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an IP Access rule

`DELETE /accounts/{account_id}/firewall/access_rules/rules/{rule_id}`

Operation ID: `ip-access-rules-for-an-account-delete-an-ip-access-rule`

Deletes an existing IP Access rule defined at the account level.

Note: This operation will affect all zones in the account.

## Definition

```yaml
{"operationId": "ip-access-rules-for-an-account-delete-an-ip-access-rule", "summary": "Delete an IP Access rule", "description": "Deletes an existing IP Access rule defined at the account level.\n\nNote: This operation will affect all zones in the account.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rule_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_api-response-single-id"}}}}, "4XX": {"description": "Delete an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_api-response-single-id"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for an account"], "x-api-token-group": ["Account Firewall Access Rules Write"], "x-cfPermissionsRequired": {"enum": ["#waf:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.access-rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
