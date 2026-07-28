---
title: List Vectors
page_id: operation-get-accounts-account-id-vectorize-v2-indexes-index-name-list-1480d9dc
path: operations/vectorize
description: Returns a paginated list of vector identifiers from the specified index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/list
operation_ids:
    - vectorize-list-vectors
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Vectors

`GET /accounts/{account_id}/vectorize/v2/indexes/{index_name}/list`

Operation ID: `vectorize-list-vectors`

Returns a paginated list of vector identifiers from the specified index.

## Definition

```yaml
{"operationId": "vectorize-list-vectors", "summary": "List Vectors", "description": "Returns a paginated list of vector identifiers from the specified index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}, {"name": "count", "in": "query", "schema": {"description": "Maximum number of vectors to return", "type": "integer", "example": 50, "default": 100, "maximum": 1000, "minimum": 1}}, {"name": "cursor", "in": "query", "schema": {"description": "Cursor for pagination to get the next page of results", "type": "string", "example": "suUTaDY5PFUiRweVccnzyt9n75suNPbXHPshvCzue5mHjtj7Letjvzlza9eGj099"}}], "responses": {"200": {"description": "List Vectors Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-list-vectors-response"}}}]}}}}, "4XX": {"description": "List Vectors Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.read"]}}
```
