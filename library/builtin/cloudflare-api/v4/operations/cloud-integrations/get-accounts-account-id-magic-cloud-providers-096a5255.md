---
title: List Cloud Integrations
page_id: operation-get-accounts-account-id-magic-cloud-providers-3fc7ec88
path: operations/cloud-integrations
description: List Cloud Integrations (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/providers
operation_ids:
    - providers-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Cloud Integrations

`GET /accounts/{account_id}/magic/cloud/providers`

Operation ID: `providers-list`

List Cloud Integrations (Closed Beta).

## Definition

```yaml
{"operationId": "providers-list", "summary": "List Cloud Integrations", "description": "List Cloud Integrations (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "status", "in": "query", "schema": {"type": "boolean"}}, {"name": "order_by", "in": "query", "description": "One of [\"updated_at\", \"id\", \"cloud_type\", \"name\"].", "schema": {"type": "string"}}, {"name": "desc", "in": "query", "schema": {"type": "boolean"}}, {"name": "cloudflare", "in": "query", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_read_account_providers_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Cloud Integrations"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
