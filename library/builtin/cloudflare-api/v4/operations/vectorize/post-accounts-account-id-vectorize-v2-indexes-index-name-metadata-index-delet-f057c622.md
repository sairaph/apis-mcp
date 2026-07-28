---
title: Delete Metadata Index
page_id: operation-post-accounts-account-id-vectorize-v2-indexes-index-name-metadata-index-03a12ddb
path: operations/vectorize
description: Allow Vectorize to delete the specified metadata index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/delete
operation_ids:
    - vectorize-delete-metadata-index
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Metadata Index

`POST /accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/delete`

Operation ID: `vectorize-delete-metadata-index`

Allow Vectorize to delete the specified metadata index.

## Definition

```yaml
{"operationId": "vectorize-delete-metadata-index", "summary": "Delete Metadata Index", "description": "Allow Vectorize to delete the specified metadata index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_delete-metadata-index-request"}}}}, "responses": {"200": {"description": "Delete Metadata Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_delete-metadata-index-response"}}}]}}}}, "4XX": {"description": "Delete Metadata Index Failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.delete"]}}
```
