---
title: Upsert Vectors (Deprecated)
page_id: operation-post-accounts-account-id-vectorize-indexes-index-name-upsert-86b2b63c
path: operations/vectorize-beta-deprecated
description: Upserts vectors into the specified index, creating them if they do not exist and returns the count of values and ids successfully inserted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/indexes/{index_name}/upsert
operation_ids:
    - vectorize-(-deprecated)-upsert-vector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upsert Vectors (Deprecated)

`POST /accounts/{account_id}/vectorize/indexes/{index_name}/upsert`

Operation ID: `vectorize-(-deprecated)-upsert-vector`

Upserts vectors into the specified index, creating them if they do not exist and returns the count of values and ids successfully inserted.

## Definition

```yaml
{"operationId": "vectorize-(-deprecated)-upsert-vector", "summary": "Upsert Vectors (Deprecated)", "description": "Upserts vectors into the specified index, creating them if they do not exist and returns the count of values and ids successfully inserted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/x-ndjson": {"schema": {"description": "ndjson file containing vectors to upsert.", "type": "string", "format": "binary", "example": "@/path/to/vectors.ndjson"}}}}, "responses": {"200": {"description": "Insert Vectors Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-upsert-response"}}}]}}}}, "4XX": {"description": "Insert Vectors Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize Beta (Deprecated)"], "x-api-token-group": ["Vectorize Write"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of the POST `/accounts/{account_id}/vectorize/v2/indexes/{index_name}/upsert` endpoint.", "display": true, "id": "vectorize_upsert_deprecation"}}
```
