---
title: Validate origin
page_id: operation-post-accounts-account-id-logpush-validate-origin-aa186c90
path: operations/logpush-jobs-for-an-account
description: Validates logpull origin with logpull_options.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/logpush/validate/origin
operation_ids:
    - post-accounts-account_id-logpush-validate-origin
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate origin

`POST /accounts/{account_id}/logpush/validate/origin`

Operation ID: `post-accounts-account_id-logpush-validate-origin`

Validates logpull origin with logpull_options.

## Definition

```yaml
{"operationId": "post-accounts-account_id-logpush-validate-origin", "summary": "Validate origin", "description": "Validates logpull origin with logpull_options.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"logpull_options": {"$ref": "#/components/schemas/logpush_logpull_options"}}, "required": ["logpull_options"]}}}}, "responses": {"200": {"description": "Validate origin response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_validate_response"}}}}, "4XX": {"description": "Validate origin response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for an account"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.account-validate.origin", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
