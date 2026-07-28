---
title: Validate rule with context
page_id: operation-post-accounts-account-id-cloudforce-one-rules-validate-7613c0dd
path: operations/rules
description: Validate rule syntax, name uniqueness, namespace, and meta checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/rules/validate
operation_ids:
    - cloudforce-one-validate-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate rule with context

`POST /accounts/{account_id}/cloudforce-one/rules/validate`

Operation ID: `cloudforce-one-validate-rule`

Validate rule syntax, name uniqueness, namespace, and meta checks.

## Definition

```yaml
{"operationId": "cloudforce-one-validate-rule", "summary": "Validate rule with context", "description": "Validate rule syntax, name uniqueness, namespace, and meta checks.", "parameters": [{"$ref": "#/components/parameters/cloudforce-one_account_id"}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"content": {"minLength": 1, "type": "string"}, "excludeRuleId": {"type": "string", "format": "uuid"}, "name": {"type": "string", "maxLength": 255, "minLength": 1}, "namespaces": {"type": "array", "items": {"maxLength": 255, "minLength": 1, "type": "string"}, "default": []}, "path": {"type": "string", "minLength": 1}}, "required": ["name", "content"]}}}}, "responses": {"200": {"description": "Validation result.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ValidationResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Rules"]}
```
