---
title: Evaluate flag
page_id: operation-get-accounts-account-id-flagship-apps-app-id-evaluate-378dfce1
path: operations/evaluation
description: Evaluates a flag against the provided context. Pass context attributes as query parameters; values are forwarded as strings. For low-latency in-Worker evaluation, prefer the Flagship binding over this endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/flagship/apps/{app_id}/evaluate
operation_ids:
    - flagship_evaluate_flag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Evaluate flag

`GET /accounts/{account_id}/flagship/apps/{app_id}/evaluate`

Operation ID: `flagship_evaluate_flag`

Evaluates a flag against the provided context. Pass context attributes as query parameters; values are forwarded as strings. For low-latency in-Worker evaluation, prefer the Flagship binding over this endpoint.

## Definition

```yaml
{"operationId": "flagship_evaluate_flag", "summary": "Evaluate flag", "description": "Evaluates a flag against the provided context. Pass context attributes as query parameters; values are forwarded as strings. For low-latency in-Worker evaluation, prefer the Flagship binding over this endpoint.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"description": "Cloudflare account ID.", "type": "string"}}, {"name": "app_id", "in": "path", "description": "App identifier.", "required": true, "schema": {"description": "App identifier.", "type": "string"}}, {"name": "flagKey", "in": "query", "description": "The flag key to evaluate.", "required": true, "schema": {"description": "The flag key to evaluate.", "type": "string"}}, {"name": "targetingKey", "in": "query", "description": "Context targeting key (per OpenFeature spec); used for percentage rollout bucketing.", "schema": {"description": "Context targeting key (per OpenFeature spec); used for percentage rollout bucketing.", "type": "string"}}], "responses": {"200": {"description": "Flag evaluation result.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_EvaluationResult"}}}}, "400": {"description": "Missing or invalid flagKey.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "404": {"description": "Flag or app not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "500": {"description": "Evaluation error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "503": {"description": "Flag store temporarily unavailable.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Evaluation"], "x-api-token-group": ["Flagship Read", "Flagship Write", "Flagship Evaluate"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.flagship.app.evaluate"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "flagship.apps.evaluate", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
