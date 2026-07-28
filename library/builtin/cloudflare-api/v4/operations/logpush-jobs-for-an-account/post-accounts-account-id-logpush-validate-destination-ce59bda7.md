---
title: Validate destination
page_id: operation-post-accounts-account-id-logpush-validate-destination-2c59235b
path: operations/logpush-jobs-for-an-account
description: Validates destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/logpush/validate/destination
operation_ids:
    - delete-accounts-account_id-logpush-validate-destination
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Validate destination

`POST /accounts/{account_id}/logpush/validate/destination`

Operation ID: `delete-accounts-account_id-logpush-validate-destination`

Validates destination.

## Definition

```yaml
{"operationId": "delete-accounts-account_id-logpush-validate-destination", "summary": "Validate destination", "description": "Validates destination.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/logpush_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"destination_conf": {"$ref": "#/components/schemas/logpush_destination_conf"}}, "required": ["destination_conf"]}}}}, "responses": {"200": {"description": "Validate destination response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/logpush_validate_response"}}}}, "4XX": {"description": "Validate destination response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Logpush jobs for an account"], "x-api-token-group": ["Logs Write"], "x-cfPermissionsRequired": {"enum": ["#logs:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logpush.account-validate.destination", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
