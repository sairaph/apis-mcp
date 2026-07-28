---
title: Update an IP Access rule
page_id: operation-patch-user-firewall-access-rules-rules-rule-id-8da00902
path: operations/ip-access-rules-for-a-user
description: Updates an IP Access rule defined at the user level. You can only update the rule action (`mode` parameter) and notes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /user/firewall/access_rules/rules/{rule_id}
operation_ids:
    - ip-access-rules-for-a-user-update-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an IP Access rule

`PATCH /user/firewall/access_rules/rules/{rule_id}`

Operation ID: `ip-access-rules-for-a-user-update-an-ip-access-rule`

Updates an IP Access rule defined at the user level. You can only update the rule action (`mode` parameter) and notes.

## Definition

```yaml
{"operationId": "ip-access-rules-for-a-user-update-an-ip-access-rule", "summary": "Update an IP Access rule", "description": "Updates an IP Access rule defined at the user level. You can only update the rule action (`mode` parameter) and notes.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rule_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"mode": {"$ref": "#/components/schemas/firewall_schemas-mode"}, "notes": {"$ref": "#/components/schemas/firewall_notes"}}}}}}, "responses": {"200": {"description": "Update an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_rule_single_response"}}}}, "4XX": {"description": "Update an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_rule_single_response"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for a user"], "x-api-token-group": ["Account Firewall Access Rules Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.firewall.access-rules", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
