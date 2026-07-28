---
title: Create Vectorize Index
page_id: operation-post-accounts-account-id-vectorize-v2-indexes-5ef1ec59
path: operations/vectorize
description: Creates and returns a new Vectorize Index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes
operation_ids:
    - vectorize-create-vectorize-index
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Vectorize Index

`POST /accounts/{account_id}/vectorize/v2/indexes`

Operation ID: `vectorize-create-vectorize-index`

Creates and returns a new Vectorize Index.

## Definition

```yaml
{"operationId": "vectorize-create-vectorize-index", "summary": "Create Vectorize Index", "description": "Creates and returns a new Vectorize Index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_create-index-request"}}}}, "responses": {"200": {"description": "Create Vectorize Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_create-index-response"}}}]}}}}, "4XX": {"description": "Create Vectorize Index Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.create"]}}
```
