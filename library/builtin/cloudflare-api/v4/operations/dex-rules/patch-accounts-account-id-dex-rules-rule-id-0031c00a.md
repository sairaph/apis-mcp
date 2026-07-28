---
title: Update a DEX Rule
page_id: operation-patch-accounts-account-id-dex-rules-rule-id-1617faa0
path: operations/dex-rules
description: Update a DEX Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/dex/rules/{rule_id}
operation_ids:
    - update-dex-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a DEX Rule

`PATCH /accounts/{account_id}/dex/rules/{rule_id}`

Operation ID: `update-dex-rule`

Update a DEX Rule.

## Definition

```yaml
{"operationId": "update-dex-rule", "summary": "Update a DEX Rule", "description": "Update a DEX Rule.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "rule_id", "in": "path", "description": "Unique identifier of the rule.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_patch_rule_body"}}}}, "responses": {"200": {"description": "Success response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-rule"}}}]}}}}, "4XX": {"description": "Update DEX Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Rules"], "x-api-token-group": ["Cloudflare DEX Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.rules", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
