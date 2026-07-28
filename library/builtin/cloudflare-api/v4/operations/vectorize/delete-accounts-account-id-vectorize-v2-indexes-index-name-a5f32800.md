---
title: Delete Vectorize Index
page_id: operation-delete-accounts-account-id-vectorize-v2-indexes-index-name-ce724257
path: operations/vectorize
description: Deletes the specified Vectorize Index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}
operation_ids:
    - vectorize-delete-vectorize-index
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Vectorize Index

`DELETE /accounts/{account_id}/vectorize/v2/indexes/{index_name}`

Operation ID: `vectorize-delete-vectorize-index`

Deletes the specified Vectorize Index.

## Definition

```yaml
{"operationId": "vectorize-delete-vectorize-index", "summary": "Delete Vectorize Index", "description": "Deletes the specified Vectorize Index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "responses": {"200": {"description": "Delete Vectorize Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}}]}}}}, "4XX": {"description": "Delete Vectorize Index Failure", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.delete"]}}
```
