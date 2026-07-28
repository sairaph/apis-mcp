---
title: Get an IP Access rule
page_id: operation-get-user-firewall-access-rules-rules-rule-id-daae6944
path: operations/ip-access-rules-for-a-user
description: Fetches the details of an IP Access rule defined at the user level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/firewall/access_rules/rules/{rule_id}
operation_ids:
    - ip-access-rules-for-a-user-get-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an IP Access rule

`GET /user/firewall/access_rules/rules/{rule_id}`

Operation ID: `ip-access-rules-for-a-user-get-an-ip-access-rule`

Fetches the details of an IP Access rule defined at the user level.

## Definition

```yaml
{"operationId": "ip-access-rules-for-a-user-get-an-ip-access-rule", "summary": "Get an IP Access rule", "description": "Fetches the details of an IP Access rule defined at the user level.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rule_identifier"}}], "responses": {"200": {"description": "Get an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_rule_single_response"}}}}, "4XX": {"description": "Get an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_rule_single_response"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for a user"], "x-api-token-group": ["Account Firewall Access Rules Write", "Account Firewall Access Rules Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
