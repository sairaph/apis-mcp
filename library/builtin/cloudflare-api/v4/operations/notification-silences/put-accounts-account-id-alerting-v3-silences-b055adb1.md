---
title: Update Silences
page_id: operation-put-accounts-account-id-alerting-v3-silences-013804f5
path: operations/notification-silences
description: Updates existing silences for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/alerting/v3/silences
operation_ids:
    - notification-silences-update-silences
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Silences

`PUT /accounts/{account_id}/alerting/v3/silences`

Operation ID: `notification-silences-update-silences`

Updates existing silences for an account.

## Definition

```yaml
{"operationId": "notification-silences-update-silences", "summary": "Update Silences", "description": "Updates existing silences for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/aaa_silence_update_request"}}}}}, "responses": {"200": {"description": "Update Silences response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_silences_components-schemas-response_collection"}}}}, "4XX": {"description": "Update Silences response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_silences_components-schemas-response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification Silences"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
