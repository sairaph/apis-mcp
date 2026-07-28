---
title: Create Silences
page_id: operation-post-accounts-account-id-alerting-v3-silences-64d8c022
path: operations/notification-silences
description: Creates a new silence for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/alerting/v3/silences
operation_ids:
    - notification-silences-create-silences
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Silences

`POST /accounts/{account_id}/alerting/v3/silences`

Operation ID: `notification-silences-create-silences`

Creates a new silence for an account.

## Definition

```yaml
{"operationId": "notification-silences-create-silences", "summary": "Create Silences", "description": "Creates a new silence for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/aaa_silence_create_request"}}}}}, "responses": {"200": {"description": "Create Silences response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-2"}}}}, "4XX": {"description": "Create Silences response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_silences_components-schemas-response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification Silences"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
