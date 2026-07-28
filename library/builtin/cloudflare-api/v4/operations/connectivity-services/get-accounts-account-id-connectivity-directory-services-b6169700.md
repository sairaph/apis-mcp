---
title: List Workers VPC connectivity services
page_id: operation-get-accounts-account-id-connectivity-directory-services-82f961a6
path: operations/connectivity-services
description: List Workers VPC connectivity services
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/connectivity/directory/services
operation_ids:
    - connectivity-services-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Workers VPC connectivity services

`GET /accounts/{account_id}/connectivity/directory/services`

Operation ID: `connectivity-services-list`

## Definition

```yaml
{"operationId": "connectivity-services-list", "summary": "List Workers VPC connectivity services", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}, {"name": "type", "in": "query", "schema": {"type": "string", "nullable": true, "oneOf": [{"$ref": "#/components/schemas/infra_ServiceType"}]}}, {"name": "page", "in": "query", "description": "Current page in the response", "schema": {"type": "integer", "format": "int32", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Max amount of entries returned per page", "schema": {"type": "integer", "format": "int32", "default": 1000, "maximum": 1000, "minimum": 1}}], "responses": {"200": {"description": "Successfully retrieved Workers VPC connectivity services", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_ConnectivityServiceArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to retrieve Workers VPC connectivity services", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Connectivity Services"]}
```
