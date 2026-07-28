---
title: Evaluate flag (POST)
page_id: operation-post-accounts-account-id-flagship-apps-app-id-evaluate-ca2e08ad
path: operations/evaluation
description: Evaluates a flag against the provided context, passed as a JSON request body (OFREP-shaped) rather than query parameters. Returns the same response shape as the GET variant.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/flagship/apps/{app_id}/evaluate
operation_ids:
    - flagship_evaluate_flag_post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Evaluate flag (POST)

`POST /accounts/{account_id}/flagship/apps/{app_id}/evaluate`

Operation ID: `flagship_evaluate_flag_post`

Evaluates a flag against the provided context, passed as a JSON request body (OFREP-shaped) rather than query parameters. Returns the same response shape as the GET variant.

## Definition

```yaml
{"operationId": "flagship_evaluate_flag_post", "summary": "Evaluate flag (POST)", "description": "Evaluates a flag against the provided context, passed as a JSON request body (OFREP-shaped) rather than query parameters. Returns the same response shape as the GET variant.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"description": "Cloudflare account ID.", "type": "string"}}, {"name": "app_id", "in": "path", "description": "App identifier.", "required": true, "schema": {"description": "App identifier.", "type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"context": {"type": "object", "default": {}, "additionalProperties": {"$ref": "#/components/schemas/flagship_JsonValue"}}, "flagKey": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_-]+$"}}, "additionalProperties": false, "required": ["flagKey"]}}}}, "responses": {"200": {"description": "Flag evaluation result.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_EvaluationResult"}}}}, "400": {"description": "Missing or invalid flagKey, or malformed context.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "404": {"description": "Flag or app not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "500": {"description": "Evaluation error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "503": {"description": "Flag store temporarily unavailable.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Evaluation"], "x-api-token-group": ["Flagship Read", "Flagship Write", "Flagship Evaluate"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.flagship.app.evaluate"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
