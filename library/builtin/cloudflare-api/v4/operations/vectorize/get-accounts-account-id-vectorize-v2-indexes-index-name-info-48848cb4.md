---
title: Get Vectorize Index Info
page_id: operation-get-accounts-account-id-vectorize-v2-indexes-index-name-info-8c8df184
path: operations/vectorize
description: Get information about a vectorize index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/info
operation_ids:
    - vectorize-index-info
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Vectorize Index Info

`GET /accounts/{account_id}/vectorize/v2/indexes/{index_name}/info`

Operation ID: `vectorize-index-info`

Get information about a vectorize index.

## Definition

```yaml
{"operationId": "vectorize-index-info", "summary": "Get Vectorize Index Info", "description": "Get information about a vectorize index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "responses": {"200": {"description": "Get Vectorize Index Info Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-info-response"}}}]}}}}, "4XX": {"description": "Get Vectorize Index Info Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.read"]}}
```
