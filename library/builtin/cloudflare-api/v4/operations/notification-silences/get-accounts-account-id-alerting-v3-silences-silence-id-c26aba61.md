---
title: Get Silence
page_id: operation-get-accounts-account-id-alerting-v3-silences-silence-id-affbaa5f
path: operations/notification-silences
description: Gets a specific silence for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/silences/{silence_id}
operation_ids:
    - notification-silences-get-silence
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Silence

`GET /accounts/{account_id}/alerting/v3/silences/{silence_id}`

Operation ID: `notification-silences-get-silence`

Gets a specific silence for an account.

## Definition

```yaml
{"operationId": "notification-silences-get-silence", "summary": "Get Silence", "description": "Gets a specific silence for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "silence_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_silence-id"}}], "responses": {"200": {"description": "Get Silence response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_silence_components-schemas-response_collection"}}}}, "4XX": {"description": "Get Silence response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_silence_components-schemas-response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification Silences"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"]}
```
