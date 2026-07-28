---
title: Get Workers VPC connectivity service
page_id: operation-get-accounts-account-id-connectivity-directory-services-service-id-782ac627
path: operations/connectivity-services
description: Get Workers VPC connectivity service
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/connectivity/directory/services/{service_id}
operation_ids:
    - connectivity-services-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Workers VPC connectivity service

`GET /accounts/{account_id}/connectivity/directory/services/{service_id}`

Operation ID: `connectivity-services-get`

## Definition

```yaml
{"operationId": "connectivity-services-get", "summary": "Get Workers VPC connectivity service", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "service_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Successfully retrieved Workers VPC connectivity service", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_ServiceConfig"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to retrieve Workers VPC connectivity service", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Connectivity Services"]}
```
