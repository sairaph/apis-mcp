---
title: List Metadata Indexes
page_id: operation-get-accounts-account-id-vectorize-v2-indexes-index-name-metadata-index-l-8842ef47
path: operations/vectorize
description: List Metadata Indexes for the specified Vectorize Index.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/list
operation_ids:
    - vectorize-list-metadata-indexes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Metadata Indexes

`GET /accounts/{account_id}/vectorize/v2/indexes/{index_name}/metadata_index/list`

Operation ID: `vectorize-list-metadata-indexes`

List Metadata Indexes for the specified Vectorize Index.

## Definition

```yaml
{"operationId": "vectorize-list-metadata-indexes", "summary": "List Metadata Indexes", "description": "List Metadata Indexes for the specified Vectorize Index.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_identifier"}}, {"name": "index_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/vectorize_index-name"}}], "responses": {"200": {"description": "List Metadata Index Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/vectorize_list-metadata-index-response"}}}]}}}}, "4XX": {"description": "List Metadata Index Failure Response", "content": {"application/json": {"schema": {"allOf": [{"allOf": [{"$ref": "#/components/schemas/vectorize_api-response-single"}, {"properties": {"result": {"type": "object", "nullable": true}}, "type": "object"}]}, {"$ref": "#/components/schemas/vectorize_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Vectorize"], "x-api-token-group": ["Vectorize Write", "Vectorize Read"]}
```
