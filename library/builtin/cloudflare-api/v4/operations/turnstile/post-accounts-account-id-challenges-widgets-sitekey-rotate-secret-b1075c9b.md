---
title: Rotate Secret for a Turnstile Widget
page_id: operation-post-accounts-account-id-challenges-widgets-sitekey-rotate-secret-8208ab7a
path: operations/turnstile
description: |-
    Generate a new secret key for this widget. If `invalidate_immediately`
    is set to `false`, the previous secret remains valid for 2 hours.

    Note that secrets cannot be rotated again during the grace period.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/challenges/widgets/{sitekey}/rotate_secret
operation_ids:
    - accounts-turnstile-widget-rotate-secret
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rotate Secret for a Turnstile Widget

`POST /accounts/{account_id}/challenges/widgets/{sitekey}/rotate_secret`

Operation ID: `accounts-turnstile-widget-rotate-secret`

Generate a new secret key for this widget. If `invalidate_immediately`
is set to `false`, the previous secret remains valid for 2 hours.

Note that secrets cannot be rotated again during the grace period.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_identifier"}}, {"name": "sitekey", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_sitekey"}}]
```

## Definition

```yaml
{"operationId": "accounts-turnstile-widget-rotate-secret", "summary": "Rotate Secret for a Turnstile Widget", "description": "Generate a new secret key for this widget. If `invalidate_immediately`\nis set to `false`, the previous secret remains valid for 2 hours.\n\nNote that secrets cannot be rotated again during the grace period.\n", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"invalidate_immediately": {"$ref": "#/components/schemas/turnstile_invalidate_immediately"}}}}}}, "responses": {"200": {"description": "Rotate Secret Response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/turnstile_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/turnstile_widget_detail"}}, "type": "object"}]}}}}, "4XX": {"description": "Rotate Secret Response Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/turnstile_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Turnstile"], "x-api-token-group": ["Turnstile Sites Write", "Account Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "turnstile.widgets", "x-fern-sdk-method-name": "rotate-secret", "x-forge-hidden": true}
```
