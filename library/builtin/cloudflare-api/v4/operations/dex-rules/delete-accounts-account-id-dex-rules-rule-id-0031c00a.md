---
title: Delete a DEX Rule
page_id: operation-delete-accounts-account-id-dex-rules-rule-id-38ec97c1
path: operations/dex-rules
description: Delete a DEX Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dex/rules/{rule_id}
operation_ids:
    - delete-dex-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a DEX Rule

`DELETE /accounts/{account_id}/dex/rules/{rule_id}`

Operation ID: `delete-dex-rule`

Delete a DEX Rule.

## Definition

```yaml
{"operationId": "delete-dex-rule", "summary": "Delete a DEX Rule", "description": "Delete a DEX Rule.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "rule_id", "in": "path", "description": "Unique identifier of the rule.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}], "responses": {"200": {"description": "Success response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"type": "boolean", "nullable": true}}}]}}}}, "4XX": {"description": "Update DEX Rule failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Rules"], "x-api-token-group": ["Cloudflare DEX Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
