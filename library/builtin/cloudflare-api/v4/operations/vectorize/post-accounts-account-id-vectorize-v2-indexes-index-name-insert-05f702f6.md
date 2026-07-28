---
title: Insert Vectors
page_id: operation-post-accounts-account-id-vectorize-v2-indexes-index-name-insert-3d246c67
path: operations/vectorize
description: Inserts vectors into the specified index and returns a mutation id corresponding to the vectors enqueued for insertion.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/insert
operation_ids:
    - vectorize-insert-vector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Insert Vectors

`POST /accounts/{account_id}/vectorize/v2/indexes/{index_name}/insert`

Operation ID: `vectorize-insert-vector`

Inserts vectors into the specified index and returns a mutation id corresponding to the vectors enqueued for insertion.

## Definition

```yaml
{"operationId": "vectorize-insert-vector", "summary": "Insert Vectors", "description": "Inserts vectors into the specified index and returns a mutation id corresponding to the vectors enqueued for insertion.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}, {"name": "unparsable-behavior", "in": "query", "schema": {"description": "Behavior for ndjson parse failures.", "type": "string", "enum": ["error", "discard"]}}], "requestBody": {"required": true, "content": {"application/x-ndjson": {"schema": {"description": "ndjson file containing vectors to insert.", "type": "string", "format": "binary", "example": "@/path/to/vectors.ndjson"}}}}, "responses": {"200": {"description": "Insert Vectors Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-insert-v2-response"}}}]}}}}, "4XX": {"description": "Insert Vectors Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write"]}
```
