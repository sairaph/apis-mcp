---
title: Delete Vectors By Identifier (Deprecated)
page_id: operation-post-accounts-account-id-vectorize-indexes-index-name-delete-by-ids-bc9a6f8e
path: operations/vectorize-beta-deprecated
description: Delete a set of vectors from an index by their vector identifiers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/indexes/{index_name}/delete-by-ids
operation_ids:
    - vectorize-(-deprecated)-delete-vectors-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Vectors By Identifier (Deprecated)

`POST /accounts/{account_id}/vectorize/indexes/{index_name}/delete-by-ids`

Operation ID: `vectorize-(-deprecated)-delete-vectors-by-id`

Delete a set of vectors from an index by their vector identifiers.

## Definition

```yaml
{"operationId": "vectorize-(-deprecated)-delete-vectors-by-id", "summary": "Delete Vectors By Identifier (Deprecated)", "description": "Delete a set of vectors from an index by their vector identifiers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_index-delete-vectors-by-id-request"}}}}, "responses": {"200": {"description": "Delete Vector Identifiers Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-delete-vectors-by-id-response"}}}]}}}}, "4XX": {"description": "Delete Vector Identifiers Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize Beta (Deprecated)"], "x-api-token-group": ["Vectorize Write"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of the POST `/accounts/{account_id}/vectorize/v2/indexes/{index_name}/delete_by_ids` endpoint.", "display": true, "id": "vectorize_delete_by_ids_deprecation"}}
```
