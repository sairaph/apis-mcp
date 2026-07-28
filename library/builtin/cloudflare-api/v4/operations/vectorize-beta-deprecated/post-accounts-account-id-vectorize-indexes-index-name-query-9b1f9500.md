---
title: Query Vectors (Deprecated)
page_id: operation-post-accounts-account-id-vectorize-indexes-index-name-query-d667809d
path: operations/vectorize-beta-deprecated
description: Finds vectors closest to a given vector in an index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/indexes/{index_name}/query
operation_ids:
    - vectorize-(-deprecated)-query-vector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Query Vectors (Deprecated)

`POST /accounts/{account_id}/vectorize/indexes/{index_name}/query`

Operation ID: `vectorize-(-deprecated)-query-vector`

Finds vectors closest to a given vector in an index.

## Definition

```yaml
{"operationId": "vectorize-(-deprecated)-query-vector", "summary": "Query Vectors (Deprecated)", "description": "Finds vectors closest to a given vector in an index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_index-query-request"}}}}, "responses": {"200": {"description": "Query Vectors Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-query-response"}}}]}}}}, "4XX": {"description": "Query Vectors Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize Beta (Deprecated)"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of the POST `/accounts/{account_id}/vectorize/v2/indexes/{index_name}/query` endpoint.", "display": true, "id": "vectorize_query_deprecation"}, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.read"]}}
```
