---
title: Get managed rules
page_id: operation-get-accounts-account-id-cloudforce-one-rules-managed-560d3342
path: operations/rules
description: Get DFP managed rule metadata (name and description) from YARA rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/managed
operation_ids:
    - cloudforce-one-get-managed-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get managed rules

`GET /accounts/{account_id}/cloudforce-one/rules/managed`

Operation ID: `cloudforce-one-get-managed-rules`

Get DFP managed rule metadata (name and description) from YARA rules.

## Definition

```yaml
{"operationId": "cloudforce-one-get-managed-rules", "summary": "Get managed rules", "description": "Get DFP managed rule metadata (name and description) from YARA rules.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "responses": {"200": {"description": "Managed rules metadata.", "content": {"application/json": {"schema": {"type": "object", "properties": {"metadata": {"type": "object", "properties": {"fetched_at": {"type": "string"}, "total_rules": {"type": "number"}}, "required": ["total_rules", "fetched_at"]}, "rules": {"type": "array", "items": {"properties": {"description": {"type": "string"}, "name": {"type": "string"}}, "required": ["name", "description"], "type": "object"}}}, "required": ["rules", "metadata"]}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
