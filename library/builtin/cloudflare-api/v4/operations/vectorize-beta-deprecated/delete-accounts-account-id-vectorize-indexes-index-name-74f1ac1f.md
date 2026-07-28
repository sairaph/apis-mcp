---
title: Delete Vectorize Index (Deprecated)
page_id: operation-delete-accounts-account-id-vectorize-indexes-index-name-bb121c4e
path: operations/vectorize-beta-deprecated
description: Deletes the specified Vectorize Index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/vectorize/indexes/{index_name}
operation_ids:
    - vectorize-(-deprecated)-delete-vectorize-index
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Vectorize Index (Deprecated)

`DELETE /accounts/{account_id}/vectorize/indexes/{index_name}`

Operation ID: `vectorize-(-deprecated)-delete-vectorize-index`

Deletes the specified Vectorize Index.

## Definition

```yaml
{"operationId": "vectorize-(-deprecated)-delete-vectorize-index", "summary": "Delete Vectorize Index (Deprecated)", "description": "Deletes the specified Vectorize Index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "responses": {"200": {"description": "Delete Vectorize Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}}]}}}}, "4XX": {"description": "Delete Vectorize Index Failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize Beta (Deprecated)"], "x-api-token-group": ["Vectorize Write"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of the DELETE `/accounts/{account_id}/vectorize/v2/indexes/{index_name}` endpoint.", "display": true, "id": "vectorize_delete_index_deprecation"}, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.delete"]}}
```
