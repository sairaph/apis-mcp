---
title: Get browser extension configuration
page_id: operation-get-accounts-account-id-browser-extension-config-196d38d3
path: operations/browser-extension-config
description: |-
    Returns the browser extension configuration for an account.

    This endpoint is currently backed by the scaffolded service response while
    the runtime configuration surface is being implemented.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-extension/config
operation_ids:
    - accounts-browser-extension-config-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get browser extension configuration

`GET /accounts/{account_id}/browser-extension/config`

Operation ID: `accounts-browser-extension-config-get`

Returns the browser extension configuration for an account.

This endpoint is currently backed by the scaffolded service response while
the runtime configuration surface is being implemented.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/brex_AccountId"}]
```

## Definition

```yaml
{"operationId": "accounts-browser-extension-config-get", "summary": "Get browser extension configuration", "description": "Returns the browser extension configuration for an account.\n\nThis endpoint is currently backed by the scaffolded service response while\nthe runtime configuration surface is being implemented.\n", "responses": {"200": {"description": "Browser extension configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/brex_ConfigResponse"}}}}, "401": {"$ref": "#/components/responses/brex_Unauthorized"}, "403": {"$ref": "#/components/responses/brex_Forbidden"}, "404": {"$ref": "#/components/responses/brex_NotFound"}, "500": {"$ref": "#/components/responses/brex_InternalServerError"}}, "security": [{"api_token": []}], "tags": ["Browser Extension Config"]}
```
