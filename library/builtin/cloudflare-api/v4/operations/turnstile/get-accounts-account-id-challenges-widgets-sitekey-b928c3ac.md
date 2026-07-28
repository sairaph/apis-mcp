---
title: Turnstile Widget Details
page_id: operation-get-accounts-account-id-challenges-widgets-sitekey-62f8d1b8
path: operations/turnstile
description: Show a single challenge widget configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/challenges/widgets/{sitekey}
operation_ids:
    - accounts-turnstile-widget-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Turnstile Widget Details

`GET /accounts/{account_id}/challenges/widgets/{sitekey}`

Operation ID: `accounts-turnstile-widget-get`

Show a single challenge widget configuration.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_identifier"}}, {"name": "sitekey", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_sitekey"}}]
```

## Definition

```yaml
{"operationId": "accounts-turnstile-widget-get", "summary": "Turnstile Widget Details", "description": "Show a single challenge widget configuration.", "responses": {"200": {"description": "Turnstile Widget Details Response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/turnstile_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/turnstile_widget_detail"}}, "type": "object"}]}}}}, "4XX": {"description": "Turnstile Widget Details Response Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/turnstile_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Turnstile"], "x-api-token-group": ["Turnstile Sites Write", "Turnstile Sites Read", "Account Settings Write", "Account Settings Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "turnstile.widgets", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
