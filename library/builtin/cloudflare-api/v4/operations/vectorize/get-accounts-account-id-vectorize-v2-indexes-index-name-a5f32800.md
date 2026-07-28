---
title: Get Vectorize Index
page_id: operation-get-accounts-account-id-vectorize-v2-indexes-index-name-e1e1a4eb
path: operations/vectorize
description: Returns the specified Vectorize Index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}
operation_ids:
    - vectorize-get-vectorize-index
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Vectorize Index

`GET /accounts/{account_id}/vectorize/v2/indexes/{index_name}`

Operation ID: `vectorize-get-vectorize-index`

Returns the specified Vectorize Index.

## Definition

```yaml
{"operationId": "vectorize-get-vectorize-index", "summary": "Get Vectorize Index", "description": "Returns the specified Vectorize Index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "responses": {"200": {"description": "Get Vectorize Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_create-index-response"}}}]}}}}, "4XX": {"description": "Get Vectorize Index Failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.read"]}}
```
