---
title: List namespaces in catalog
page_id: operation-get-accounts-account-id-r2-catalog-bucket-name-namespaces-e2b95307
path: operations/namespace-management
description: |-
    Returns a list of namespaces in the specified R2 catalog.
    Supports hierarchical filtering and pagination for efficient traversal
    of large namespace hierarchies.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces
operation_ids:
    - list-namespaces
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List namespaces in catalog

`GET /accounts/{account_id}/r2-catalog/{bucket_name}/namespaces`

Operation ID: `list-namespaces`

Returns a list of namespaces in the specified R2 catalog.
Supports hierarchical filtering and pagination for efficient traversal
of large namespace hierarchies.

## Definition

```yaml
{"operationId": "list-namespaces", "summary": "List namespaces in catalog", "description": "Returns a list of namespaces in the specified R2 catalog.\nSupports hierarchical filtering and pagination for efficient traversal\nof large namespace hierarchies.\n", "parameters": [{"name": "account_id", "in": "path", "description": "Identifies the account.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_account-id"}}, {"name": "bucket_name", "in": "path", "description": "Specifies the R2 bucket name.", "required": true, "schema": {"$ref": "#/components/schemas/r2-data-catalog_bucket-name"}}, {"name": "page_token", "in": "query", "description": "Opaque pagination token from a previous response.\nUse this to fetch the next page of results.\n", "schema": {"type": "string"}, "example": "MSYxNzU5NzU1NTc4NTA0MTk0JjAxOTliOTliLTJjODgtNzNiMy04ZGJiLTQyMWUwZThmMjc1Nw"}, {"name": "page_size", "in": "query", "description": "Maximum number of namespaces to return per page.\nDefaults to 100, maximum 1000.\n", "schema": {"type": "integer", "default": 100, "maximum": 1000, "minimum": 1}, "example": 100}, {"name": "parent", "in": "query", "description": "Parent namespace to filter by. Only returns direct children of this namespace.\nFor nested namespaces, use %1F as separator (e.g., \"bronze%1Fanalytics\").\nOmit this parameter to list top-level namespaces.\n", "schema": {"type": "string"}, "example": "bronze"}, {"name": "return_uuids", "in": "query", "description": "Whether to include namespace UUIDs in the response.\nSet to true to receive the namespace_uuids array.\n", "schema": {"type": "boolean", "default": false}, "example": true}, {"name": "return_details", "in": "query", "description": "Whether to include additional metadata (timestamps).\nWhen true, response includes created_at and updated_at arrays.\n", "schema": {"type": "boolean", "default": false}, "example": true}], "responses": {"200": {"description": "List of namespaces retrieved successfully.", "content": {"application/json": {"example": {"errors": [], "messages": [], "result": {"details": [{"created_at": "2025-10-07T10:00:00Z", "namespace": ["bronze"], "namespace_uuid": "0199b999-6869-7383-bb1f-d30e059d5326", "updated_at": "2025-10-07T12:00:00Z"}, {"created_at": "2025-10-07T10:15:00Z", "namespace": ["silver"], "namespace_uuid": "0199b99b-2c88-73b3-8dbb-421e0e8f2757", "updated_at": null}, {"created_at": "2025-10-07T10:30:00Z", "namespace": ["gold"], "namespace_uuid": "0199b99c-3d99-73c4-9dcc-532f1f9g3868", "updated_at": "2025-10-07T11:00:00Z"}], "namespace_uuids": ["0199b999-6869-7383-bb1f-d30e059d5326", "0199b99b-2c88-73b3-8dbb-421e0e8f2757", "0199b99c-3d99-73c4-9dcc-532f1f9g3868"], "namespaces": [["bronze"], ["silver"], ["gold"]], "next_page_token": null}, "success": true}, "schema": {"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-data-catalog_namespace-list-response"}}, "type": "object"}]}}}}, "400": {"description": "Bad request (e.g., invalid page_size, malformed parent namespace).", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "401": {"description": "Authentication failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "404": {"description": "Catalog not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-data-catalog_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Namespace Management"], "x-api-token-group": ["Workers R2 Data Catalog Write", "Workers R2 Data Catalog Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2-data-catalog.namespaces", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
