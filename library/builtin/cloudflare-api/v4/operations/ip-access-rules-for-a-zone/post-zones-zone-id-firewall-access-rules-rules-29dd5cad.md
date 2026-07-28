---
title: Create an IP Access rule
page_id: operation-post-zones-zone-id-firewall-access-rules-rules-32170452
path: operations/ip-access-rules-for-a-zone
description: |-
    Creates a new IP Access rule for a zone.

    Note: To create an IP Access rule that applies to multiple zones, refer to [IP Access rules for a user](#ip-access-rules-for-a-user) or [IP Access rules for an account](#ip-access-rules-for-an-account) as appropriate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/firewall/access_rules/rules
operation_ids:
    - ip-access-rules-for-a-zone-create-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an IP Access rule

`POST /zones/{zone_id}/firewall/access_rules/rules`

Operation ID: `ip-access-rules-for-a-zone-create-an-ip-access-rule`

Creates a new IP Access rule for a zone.

Note: To create an IP Access rule that applies to multiple zones, refer to [IP Access rules for a user](#ip-access-rules-for-a-user) or [IP Access rules for an account](#ip-access-rules-for-an-account) as appropriate.

## Definition

```yaml
{"operationId": "ip-access-rules-for-a-zone-create-an-ip-access-rule", "summary": "Create an IP Access rule", "description": "Creates a new IP Access rule for a zone.\n\nNote: To create an IP Access rule that applies to multiple zones, refer to [IP Access rules for a user](#ip-access-rules-for-a-user) or [IP Access rules for an account](#ip-access-rules-for-an-account) as appropriate.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"configuration": {"$ref": "#/components/schemas/firewall_configuration"}, "mode": {"$ref": "#/components/schemas/firewall_schemas-mode"}, "notes": {"allOf": [{"$ref": "#/components/schemas/firewall_notes"}, {"default": ""}]}}, "required": ["mode", "configuration"]}}}}, "responses": {"200": {"description": "Create an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_rule_single_response"}}}}, "4XX": {"description": "Create an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_rule_single_response"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for a zone"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.zone-access-rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
