---
title: Create a DEX Rule
page_id: operation-post-accounts-account-id-dex-rules-881f5ec7
path: operations/dex-rules
description: Create a DEX Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dex/rules
operation_ids:
    - create-dex-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a DEX Rule

`POST /accounts/{account_id}/dex/rules`

Operation ID: `create-dex-rule`

Create a DEX Rule.

## Definition

```yaml
{"operationId": "create-dex-rule", "summary": "Create a DEX Rule", "description": "Create a DEX Rule.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_create_rule_body"}}}}, "responses": {"200": {"description": "Success response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_dex-rule"}}}]}}}}, "4XX": {"description": "Create DEX Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Rules"], "x-api-token-group": ["Cloudflare DEX Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
