---
title: Create an IP Access rule
page_id: operation-post-accounts-account-id-firewall-access-rules-rules-24ede99b
path: operations/ip-access-rules-for-an-account
description: |-
    Creates a new IP Access rule for an account. The rule will apply to all zones in the account.

    Note: To create an IP Access rule that applies to a single zone, refer to the [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/firewall/access_rules/rules
operation_ids:
    - ip-access-rules-for-an-account-create-an-ip-access-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an IP Access rule

`POST /accounts/{account_id}/firewall/access_rules/rules`

Operation ID: `ip-access-rules-for-an-account-create-an-ip-access-rule`

Creates a new IP Access rule for an account. The rule will apply to all zones in the account.

Note: To create an IP Access rule that applies to a single zone, refer to the [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.

## Definition

```yaml
{"operationId": "ip-access-rules-for-an-account-create-an-ip-access-rule", "summary": "Create an IP Access rule", "description": "Creates a new IP Access rule for an account. The rule will apply to all zones in the account.\n\nNote: To create an IP Access rule that applies to a single zone, refer to the [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"configuration": {"$ref": "#/components/schemas/firewall_configuration"}, "mode": {"$ref": "#/components/schemas/firewall_schemas-mode"}, "notes": {"allOf": [{"$ref": "#/components/schemas/firewall_notes"}, {"default": ""}]}}, "required": ["mode", "configuration"]}}}}, "responses": {"200": {"description": "Create an IP Access rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_response_single"}}}}, "4XX": {"description": "Create an IP Access rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_response_single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["IP Access rules for an account"], "x-api-token-group": ["Account Firewall Access Rules Write"], "x-cfPermissionsRequired": {"enum": ["#waf:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "firewall.access-rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
