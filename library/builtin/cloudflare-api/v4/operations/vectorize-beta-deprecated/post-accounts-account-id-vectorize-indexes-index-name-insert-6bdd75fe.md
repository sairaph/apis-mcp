---
title: Insert Vectors (Deprecated)
page_id: operation-post-accounts-account-id-vectorize-indexes-index-name-insert-a54045a1
path: operations/vectorize-beta-deprecated
description: Inserts vectors into the specified index and returns the count of the vectors successfully inserted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/indexes/{index_name}/insert
operation_ids:
    - vectorize-(-deprecated)-insert-vector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Insert Vectors (Deprecated)

`POST /accounts/{account_id}/vectorize/indexes/{index_name}/insert`

Operation ID: `vectorize-(-deprecated)-insert-vector`

Inserts vectors into the specified index and returns the count of the vectors successfully inserted.

## Definition

```yaml
{"operationId": "vectorize-(-deprecated)-insert-vector", "summary": "Insert Vectors (Deprecated)", "description": "Inserts vectors into the specified index and returns the count of the vectors successfully inserted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/x-ndjson": {"schema": {"description": "ndjson file containing vectors to insert.", "type": "string", "format": "binary", "example": "@/path/to/vectors.ndjson"}}}}, "responses": {"200": {"description": "Insert Vectors Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-insert-response"}}}]}}}}, "4XX": {"description": "Insert Vectors Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize Beta (Deprecated)"], "x-api-token-group": ["Vectorize Write"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of the POST `/accounts/{account_id}/vectorize/v2/indexes/{index_name}/insert` endpoint.", "display": true, "id": "vectorize_insert_deprecation"}, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.update"]}}
```
