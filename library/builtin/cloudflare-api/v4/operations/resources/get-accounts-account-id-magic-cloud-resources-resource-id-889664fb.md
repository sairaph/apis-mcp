---
title: Read Resource
page_id: operation-get-accounts-account-id-magic-cloud-resources-resource-id-20a0d248
path: operations/resources
description: Read an resource from the Resource Catalog (Closed Beta).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cloud/resources/{resource_id}
operation_ids:
    - resources-catalog-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Read Resource

`GET /accounts/{account_id}/magic/cloud/resources/{resource_id}`

Operation ID: `resources-catalog-read`

Read an resource from the Resource Catalog (Closed Beta).

## Definition

```yaml
{"operationId": "resources-catalog-read", "summary": "Read Resource", "description": "Read an resource from the Resource Catalog (Closed Beta).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_account_id"}}, {"name": "resource_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mcn_resource_id"}}, {"name": "v2", "in": "query", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "OK.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_read_account_resource_response"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "401": {"description": "Invalid Credentials.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "404": {"description": "Not Found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}, "500": {"description": "Internal Server Error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mcn_bad_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Resources"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read"]}
```
