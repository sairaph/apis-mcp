---
title: Delete a Turnstile Widget
page_id: operation-delete-accounts-account-id-challenges-widgets-sitekey-127f188a
path: operations/turnstile
description: Destroy a Turnstile Widget.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/challenges/widgets/{sitekey}
operation_ids:
    - accounts-turnstile-widget-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Turnstile Widget

`DELETE /accounts/{account_id}/challenges/widgets/{sitekey}`

Operation ID: `accounts-turnstile-widget-delete`

Destroy a Turnstile Widget.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_identifier"}}, {"name": "sitekey", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_sitekey"}}]
```

## Definition

```yaml
{"operationId": "accounts-turnstile-widget-delete", "summary": "Delete a Turnstile Widget", "description": "Destroy a Turnstile Widget.", "responses": {"200": {"description": "Delete Turnstile Widget Response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/turnstile_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/turnstile_widget_detail"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Turnstile Widget Response Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/turnstile_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Turnstile"], "x-api-token-group": ["Turnstile Sites Write", "Account Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "turnstile.widgets", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
