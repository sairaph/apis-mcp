---
title: Delete Vectors By Identifier
page_id: operation-post-accounts-account-id-vectorize-v2-indexes-index-name-delete-by-ids-dd769cd6
path: operations/vectorize
description: Delete a set of vectors from an index by their vector identifiers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/delete_by_ids
operation_ids:
    - vectorize-delete-vectors-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Vectors By Identifier

`POST /accounts/{account_id}/vectorize/v2/indexes/{index_name}/delete_by_ids`

Operation ID: `vectorize-delete-vectors-by-id`

Delete a set of vectors from an index by their vector identifiers.

## Definition

```yaml
{"operationId": "vectorize-delete-vectors-by-id", "summary": "Delete Vectors By Identifier", "description": "Delete a set of vectors from an index by their vector identifiers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_index-delete-vectors-by-id-request"}}}}, "responses": {"200": {"description": "Delete Vector Identifiers Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_index-delete-vectors-by-id-v2-response"}}}]}}}}, "4XX": {"description": "Delete Vector Identifiers Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.delete"]}}
```
