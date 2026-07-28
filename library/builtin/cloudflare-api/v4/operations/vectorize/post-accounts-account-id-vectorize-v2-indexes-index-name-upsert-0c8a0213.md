---
title: Upsert Vectors
page_id: operation-post-accounts-account-id-vectorize-v2-indexes-index-name-upsert-abc1ebba
path: operations/vectorize
description: Upserts vectors into the specified index, creating them if they do not exist and returns a mutation id corresponding to the vectors enqueued for upsertion.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/upsert
operation_ids:
    - vectorize-upsert-vector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upsert Vectors

`POST /accounts/{account_id}/vectorize/v2/indexes/{index_name}/upsert`

Operation ID: `vectorize-upsert-vector`

Upserts vectors into the specified index, creating them if they do not exist and returns a mutation id corresponding to the vectors enqueued for upsertion.

## Definition

```yaml
{"operationId": "vectorize-upsert-vector", "summary": "Upsert Vectors", "description": "Upserts vectors into the specified index, creating them if they do not exist and returns a mutation id corresponding to the vectors enqueued for upsertion.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}, {"name": "unparsable-behavior", "in": "query", "schema": {"description": "Behavior for ndjson parse failures.", "type": "string", "enum": ["error", "discard"]}}], "requestBody": {"required": true, "content": {"application/x-ndjson": {"schema": {"description": "ndjson file containing vectors to upsert.", "type": "string", "format": "binary", "example": "@/path/to/vectors.ndjson"}}}}, "responses": {"200": {"description": "Upsert Vectors Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-upsert-v2-response"}}}]}}}}, "4XX": {"description": "Upsert Vectors Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write"]}
```
