---
title: List DEX Rules
page_id: operation-get-accounts-account-id-dex-rules-255ae2c1
path: operations/dex-rules
description: List DEX Rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/rules
operation_ids:
    - list-dex-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List DEX Rules

`GET /accounts/{account_id}/dex/rules`

Operation ID: `list-dex-rules`

List DEX Rules.

## Definition

```yaml
{"operationId": "list-dex-rules", "summary": "List DEX Rules", "description": "List DEX Rules.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "required": true, "schema": {"type": "number", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "required": true, "schema": {"type": "number", "example": 10, "maximum": 50, "minimum": 1}}, {"name": "sort_order", "in": "query", "description": "Sort direction for sort_by property.", "schema": {"type": "string", "default": "ASC", "enum": ["ASC", "DESC"]}}, {"name": "sort_by", "in": "query", "description": "Which property to sort results by.", "schema": {"type": "string", "default": "name", "enum": ["name", "created_at", "updated_at"]}}, {"name": "name", "in": "query", "description": "Filter results by rule name.", "schema": {"type": "string"}}], "responses": {"200": {"description": "Success response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_list_rules_response"}}}]}}}}, "4XX": {"description": "List DEX Rules failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Rules"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.rules", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
