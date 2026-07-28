---
title: Update a Turnstile Widget
page_id: operation-put-accounts-account-id-challenges-widgets-sitekey-a72f202d
path: operations/turnstile
description: Update the configuration of a widget.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/challenges/widgets/{sitekey}
operation_ids:
    - accounts-turnstile-widget-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Turnstile Widget

`PUT /accounts/{account_id}/challenges/widgets/{sitekey}`

Operation ID: `accounts-turnstile-widget-update`

Update the configuration of a widget.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_identifier"}}, {"name": "sitekey", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_sitekey"}}]
```

## Definition

```yaml
{"operationId": "accounts-turnstile-widget-update", "summary": "Update a Turnstile Widget", "description": "Update the configuration of a widget.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"bot_fight_mode": {"$ref": "#/components/schemas/turnstile_bot_fight_mode"}, "clearance_level": {"$ref": "#/components/schemas/turnstile_clearance_level"}, "domains": {"$ref": "#/components/schemas/turnstile_domains"}, "ephemeral_id": {"$ref": "#/components/schemas/turnstile_ephemeral_id"}, "mode": {"$ref": "#/components/schemas/turnstile_widget_mode"}, "name": {"$ref": "#/components/schemas/turnstile_name"}, "offlabel": {"$ref": "#/components/schemas/turnstile_offlabel"}, "region": {"$ref": "#/components/schemas/turnstile_region"}}, "required": ["name", "mode", "domains"]}}}}, "responses": {"200": {"description": "Update Turnstile Widget Response", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/turnstile_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/turnstile_widget_detail"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Turnstile Widget Response Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/turnstile_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Turnstile"], "x-api-token-group": ["Turnstile Sites Write", "Account Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "turnstile.widgets", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
