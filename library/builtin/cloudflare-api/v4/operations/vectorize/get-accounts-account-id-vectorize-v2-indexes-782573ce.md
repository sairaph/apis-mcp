---
title: List Vectorize Indexes
page_id: operation-get-accounts-account-id-vectorize-v2-indexes-d59f3566
path: operations/vectorize
description: Returns a list of Vectorize Indexes
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes
operation_ids:
    - vectorize-list-vectorize-indexes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Vectorize Indexes

`GET /accounts/{account_id}/vectorize/v2/indexes`

Operation ID: `vectorize-list-vectorize-indexes`

Returns a list of Vectorize Indexes

## Definition

```yaml
{"operationId": "vectorize-list-vectorize-indexes", "summary": "List Vectorize Indexes", "description": "Returns a list of Vectorize Indexes", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}], "responses": {"200": {"description": "List Vectorize Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/vectorize_create-index-response"}}}}]}}}}, "4XX": {"description": "List Vectorize Index Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.list"]}}
```
