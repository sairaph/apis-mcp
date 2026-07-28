---
title: Delete an IP Access rule
page_id: operation-delete-user-firewall-access-rules-rules-rule-id-4b11052b
path: operations/ip-access-rules-for-a-user
description: |-
    Deletes an IP Access rule at the user level.

    Note: Deleting a user-level rule will affect all zones owned by the user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /user/firewall/access_rules/rules/{rule_id}
operation_ids:
    - ip-access-rules-for-a-user-delete-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an IP Access rule

`DELETE /user/firewall/access_rules/rules/{rule_id}`

Operation ID: `ip-access-rules-for-a-user-delete-an-ip-access-rule`

Deletes an IP Access rule at the user level.

Note: Deleting a user-level rule will affect all zones owned by the user.

## Definition

```yaml
{"operationId": "ip-access-rules-for-a-user-delete-an-ip-access-rule", "summary": "Delete an IP Access rule", "description": "Deletes an IP Access rule at the user level.\n\nNote: Deleting a user-level rule will affect all zones owned by the user.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rule_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_rule_single_id_response"}}}}, "4XX": {"description": "Delete an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_rule_single_id_response"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for a user"], "x-api-token-group": ["Account Firewall Access Rules Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.firewall.access-rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
