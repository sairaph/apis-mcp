---
title: List Silences
page_id: operation-get-accounts-account-id-alerting-v3-silences-fd7abe10
path: operations/notification-silences
description: Gets a list of silences for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/silences
operation_ids:
    - notification-silences-list-silences
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Silences

`GET /accounts/{account_id}/alerting/v3/silences`

Operation ID: `notification-silences-list-silences`

Gets a list of silences for an account.

## Definition

```yaml
{"operationId": "notification-silences-list-silences", "summary": "List Silences", "description": "Gets a list of silences for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "List Silences response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_silences_components-schemas-response_collection"}}}}, "4XX": {"description": "List Silences response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_silences_components-schemas-response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification Silences"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"]}
```
