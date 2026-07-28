---
title: Create Metadata Index
page_id: operation-post-accounts-account-id-vectorize-v2-indexes-index-name-metadata-index-8b63d7f8
path: operations/vectorize
description: Enable metadata filtering based on metadata property. Limited to 10 properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/create
operation_ids:
    - vectorize-create-metadata-index
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Metadata Index

`POST /accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/create`

Operation ID: `vectorize-create-metadata-index`

Enable metadata filtering based on metadata property. Limited to 10 properties.

## Definition

```yaml
{"operationId": "vectorize-create-metadata-index", "summary": "Create Metadata Index", "description": "Enable metadata filtering based on metadata property. Limited to 10 properties.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vectorize_create-metadata-index-request"}}}}, "responses": {"200": {"description": "Create Metadata Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_create-metadata-index-response"}}}]}}}}, "4XX": {"description": "Create Metadata Index Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.vectorize.index.create"]}}
```
