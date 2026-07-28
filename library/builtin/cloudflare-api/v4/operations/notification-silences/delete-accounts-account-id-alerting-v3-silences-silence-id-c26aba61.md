---
title: Delete Silence
page_id: operation-delete-accounts-account-id-alerting-v3-silences-silence-id-2a240d87
path: operations/notification-silences
description: Deletes an existing silence for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/alerting/v3/silences/{silence_id}
operation_ids:
    - notification-silences-delete-silences
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Silence

`DELETE /accounts/{account_id}/alerting/v3/silences/{silence_id}`

Operation ID: `notification-silences-delete-silences`

Deletes an existing silence for an account.

## Definition

```yaml
{"operationId": "notification-silences-delete-silences", "summary": "Delete Silence", "description": "Deletes an existing silence for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "silence_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_silence-id"}}], "responses": {"200": {"description": "Delete Silence response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-2"}}}}, "4XX": {"description": "Delete Silence response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_silence_components-schemas-response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification Silences"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
