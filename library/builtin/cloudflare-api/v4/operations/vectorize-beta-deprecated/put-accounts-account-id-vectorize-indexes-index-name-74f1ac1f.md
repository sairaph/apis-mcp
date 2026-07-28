---
title: Update Vectorize Index (Deprecated)
page_id: operation-put-accounts-account-id-vectorize-indexes-index-name-4c548023
path: operations/vectorize-beta-deprecated
description: Updates and returns the specified Vectorize Index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/vectorize/indexes/{index_name}
operation_ids:
    - vectorize-(-deprecated)-update-vectorize-index
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Vectorize Index (Deprecated)

`PUT /accounts/{account_id}/vectorize/indexes/{index_name}`

Operation ID: `vectorize-(-deprecated)-update-vectorize-index`

Updates and returns the specified Vectorize Index.

## Definition

```yaml
{"operationId": "vectorize-(-deprecated)-update-vectorize-index", "summary": "Update Vectorize Index (Deprecated)", "description": "Updates and returns the specified Vectorize Index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_update-index-request"}}}}, "responses": {"200": {"description": "Update Vectorize Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_create-index-response"}}}]}}}}, "4XX": {"description": "Update Vectorize Index Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize Beta (Deprecated)"], "x-api-token-group": ["Vectorize Write"], "x-cfDeprecation": {"description": "This endpoint has been deprecated and will soon be removed.", "display": true, "id": "vectorize_update_index_deprecation"}, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.update"]}}
```
