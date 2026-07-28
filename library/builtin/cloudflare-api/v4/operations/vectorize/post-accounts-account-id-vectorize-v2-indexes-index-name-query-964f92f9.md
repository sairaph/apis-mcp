---
title: Query Vectors
page_id: operation-post-accounts-account-id-vectorize-v2-indexes-index-name-query-eadc3ea8
path: operations/vectorize
description: Finds vectors closest to a given vector in an index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/query
operation_ids:
    - vectorize-query-vector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Query Vectors

`POST /accounts/{account_id}/vectorize/v2/indexes/{index_name}/query`

Operation ID: `vectorize-query-vector`

Finds vectors closest to a given vector in an index.

## Definition

```yaml
{"operationId": "vectorize-query-vector", "summary": "Query Vectors", "description": "Finds vectors closest to a given vector in an index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_index-query-v2-request"}}}}, "responses": {"200": {"description": "Query Vectors Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-query-v2-response"}}}]}}}}, "4XX": {"description": "Query Vectors Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"]}
```
