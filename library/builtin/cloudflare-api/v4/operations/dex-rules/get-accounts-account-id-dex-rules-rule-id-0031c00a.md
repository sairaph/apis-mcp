---
title: Get DEX Rule
page_id: operation-get-accounts-account-id-dex-rules-rule-id-479db437
path: operations/dex-rules
description: Get details for a DEX Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/rules/{rule_id}
operation_ids:
    - get-dex-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DEX Rule

`GET /accounts/{account_id}/dex/rules/{rule_id}`

Operation ID: `get-dex-rule`

Get details for a DEX Rule.

## Definition

```yaml
{"operationId": "get-dex-rule", "summary": "Get DEX Rule", "description": "Get details for a DEX Rule.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "rule_id", "in": "path", "description": "Unique identifier of the rule.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}], "responses": {"200": {"description": "Success response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-rule"}}}]}}}}, "4XX": {"description": "List DEX Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Rules"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.rules", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
