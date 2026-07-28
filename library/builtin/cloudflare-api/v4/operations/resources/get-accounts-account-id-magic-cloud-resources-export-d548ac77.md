---
title: Export Resources
page_id: operation-get-accounts-account-id-magic-cloud-resources-export-c294ca32
path: operations/resources
description: Export resources in the Resource Catalog as a JSON file (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/resources/export
operation_ids:
    - resources-catalog-export
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Export Resources

`GET /accounts/{account_id}/magic/cloud/resources/export`

Operation ID: `resources-catalog-export`

Export resources in the Resource Catalog as a JSON file (Closed Beta).

## Definition

```yaml
{"operationId": "resources-catalog-export", "summary": "Export Resources", "description": "Export resources in the Resource Catalog as a JSON file (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "provider_id", "in": "query", "schema": {"type": "string"}}, {"name": "resource_type", "in": "query", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_type"}}}, {"name": "resource_id", "in": "query", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}}, {"name": "region", "in": "query", "schema": {"type": "string"}}, {"name": "resource_group", "in": "query", "schema": {"type": "string"}}, {"name": "search", "in": "query", "schema": {"type": "array", "items": {"type": "string"}}}, {"name": "order_by", "in": "query", "description": "One of [\"id\", \"resource_type\", \"region\"].", "schema": {"type": "string"}}, {"name": "desc", "in": "query", "schema": {"type": "boolean"}}, {"name": "v2", "in": "query", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "Exported file.", "headers": {"Content-Disposition": {"schema": {"type": "string", "example": "attachment; filename=\"exported_resources.zip\""}}}, "content": {"application/octet-stream": {"schema": {"type": "string", "format": "binary"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Resources"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
