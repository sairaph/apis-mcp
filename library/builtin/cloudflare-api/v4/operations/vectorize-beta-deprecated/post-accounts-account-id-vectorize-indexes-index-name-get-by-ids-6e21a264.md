---
title: Get Vectors By Identifier (Deprecated)
page_id: operation-post-accounts-account-id-vectorize-indexes-index-name-get-by-ids-07a3b80e
path: operations/vectorize-beta-deprecated
description: Get a set of vectors from an index by their vector identifiers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/indexes/{index_name}/get-by-ids
operation_ids:
    - vectorize-(-deprecated)-get-vectors-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Vectors By Identifier (Deprecated)

`POST /accounts/{account_id}/vectorize/indexes/{index_name}/get-by-ids`

Operation ID: `vectorize-(-deprecated)-get-vectors-by-id`

Get a set of vectors from an index by their vector identifiers.

## Definition

```yaml
{"operationId": "vectorize-(-deprecated)-get-vectors-by-id", "summary": "Get Vectors By Identifier (Deprecated)", "description": "Get a set of vectors from an index by their vector identifiers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_index-get-vectors-by-id-request"}}}}, "responses": {"200": {"description": "Get Vectors By Identifier Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-get-vectors-by-id-response"}}}]}}}}, "4XX": {"description": "Get Vectors By Identifier Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize Beta (Deprecated)"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of the POST `/accounts/{account_id}/vectorize/v2/indexes/{index_name}/get_by_ids` endpoint.", "display": true, "id": "vectorize_get_by_ids_deprecation"}}
```
