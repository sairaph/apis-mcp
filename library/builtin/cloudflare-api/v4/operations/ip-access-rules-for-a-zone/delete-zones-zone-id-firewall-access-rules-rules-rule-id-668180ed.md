---
title: Delete an IP Access rule
page_id: operation-delete-zones-zone-id-firewall-access-rules-rules-rule-id-32354e78
path: operations/ip-access-rules-for-a-zone
description: |-
    Deletes an IP Access rule defined at the zone level.

    Optionally, you can use the `cascade` property to specify that you wish to delete similar rules in other zones managed by the same zone owner.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/firewall/access_rules/rules/{rule_id}
operation_ids:
    - ip-access-rules-for-a-zone-delete-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an IP Access rule

`DELETE /zones/{zone_id}/firewall/access_rules/rules/{rule_id}`

Operation ID: `ip-access-rules-for-a-zone-delete-an-ip-access-rule`

Deletes an IP Access rule defined at the zone level.

Optionally, you can use the `cascade` property to specify that you wish to delete similar rules in other zones managed by the same zone owner.

## Definition

```yaml
{"operationId": "ip-access-rules-for-a-zone-delete-an-ip-access-rule", "summary": "Delete an IP Access rule", "description": "Deletes an IP Access rule defined at the zone level.\n\nOptionally, you can use the `cascade` property to specify that you wish to delete similar rules in other zones managed by the same zone owner.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_rule_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"cascade": {"description": "The level to attempt to delete similar rules defined for other zones with the same owner. The default value is `none`, which will only delete the current rule. Using `basic` will delete rules that match the same action (mode) and configuration, while using `aggressive` will delete rules that match the same configuration.", "type": "string", "default": "none", "enum": ["none", "basic", "aggressive"]}}}}}}, "responses": {"200": {"description": "Delete an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_rule_single_id_response"}}}}, "4XX": {"description": "Delete an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_rule_single_id_response"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for a zone"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.zone-access-rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
