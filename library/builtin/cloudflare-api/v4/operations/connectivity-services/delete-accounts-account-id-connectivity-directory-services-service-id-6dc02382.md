---
title: Delete Workers VPC connectivity service
page_id: operation-delete-accounts-account-id-connectivity-directory-services-service-id-22c9b171
path: operations/connectivity-services
description: Delete Workers VPC connectivity service
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/connectivity/directory/services/{service_id}
operation_ids:
    - connectivity-services-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Workers VPC connectivity service

`DELETE /accounts/{account_id}/connectivity/directory/services/{service_id}`

Operation ID: `connectivity-services-delete`

## Definition

```yaml
{"operationId": "connectivity-services-delete", "summary": "Delete Workers VPC connectivity service", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "service_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Successfully deleted Workers VPC connectivity service"}, "4XX": {"description": "Failed to delete Workers VPC connectivity service", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Connectivity Services"]}
```
