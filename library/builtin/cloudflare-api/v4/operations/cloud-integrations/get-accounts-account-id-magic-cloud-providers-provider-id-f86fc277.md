---
title: Read Cloud Integration
page_id: operation-get-accounts-account-id-magic-cloud-providers-provider-id-424ee2da
path: operations/cloud-integrations
description: Read a Cloud Integration (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/providers/{provider_id}
operation_ids:
    - providers-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read Cloud Integration

`GET /accounts/{account_id}/magic/cloud/providers/{provider_id}`

Operation ID: `providers-read`

Read a Cloud Integration (Closed Beta).

## Definition

```yaml
{"operationId": "providers-read", "summary": "Read Cloud Integration", "description": "Read a Cloud Integration (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "provider_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_provider_id"}}, {"name": "status", "in": "query", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_read_account_provider_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Cloud Integrations"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
