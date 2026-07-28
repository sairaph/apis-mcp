---
title: List Resources
page_id: operation-get-accounts-account-id-magic-cloud-resources-fd922213
path: operations/resources
description: List resources in the Resource Catalog (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/resources
operation_ids:
    - resources-catalog-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Resources

`GET /accounts/{account_id}/magic/cloud/resources`

Operation ID: `resources-catalog-list`

List resources in the Resource Catalog (Closed Beta).

## Definition

```yaml
{"operationId": "resources-catalog-list", "summary": "List Resources", "description": "List resources in the Resource Catalog (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "provider_id", "in": "query", "schema": {"type": "string"}}, {"name": "resource_type", "in": "query", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_type"}}}, {"name": "resource_id", "in": "query", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}}, {"name": "region", "in": "query", "schema": {"type": "string"}}, {"name": "resource_group", "in": "query", "schema": {"type": "string"}}, {"name": "managed", "in": "query", "schema": {"type": "boolean"}}, {"name": "search", "in": "query", "schema": {"type": "array", "items": {"type": "string"}}}, {"name": "order_by", "in": "query", "description": "One of [\"id\", \"resource_type\", \"region\"].", "schema": {"type": "string"}}, {"name": "desc", "in": "query", "schema": {"type": "boolean"}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "minimum": 1}}, {"name": "page", "in": "query", "schema": {"type": "integer", "minimum": 1}}, {"name": "cloudflare", "in": "query", "schema": {"type": "boolean"}}, {"name": "v2", "in": "query", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_read_account_resources_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Resources"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
