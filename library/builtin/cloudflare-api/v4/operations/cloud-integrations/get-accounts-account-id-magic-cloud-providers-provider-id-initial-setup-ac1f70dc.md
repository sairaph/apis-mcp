---
title: Get Cloud Integration Setup Config
page_id: operation-get-accounts-account-id-magic-cloud-providers-provider-id-initial-setup-17faab22
path: operations/cloud-integrations
description: Get initial configuration to complete Cloud Integration setup (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/providers/{provider_id}/initial_setup
operation_ids:
    - providers-initial-setup
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Cloud Integration Setup Config

`GET /accounts/{account_id}/magic/cloud/providers/{provider_id}/initial_setup`

Operation ID: `providers-initial-setup`

Get initial configuration to complete Cloud Integration setup (Closed Beta).

## Definition

```yaml
{"operationId": "providers-initial-setup", "summary": "Get Cloud Integration Setup Config", "description": "Get initial configuration to complete Cloud Integration setup (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_provider_id"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_provider_initial_setup_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Cloud Integrations"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
