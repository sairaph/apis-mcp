---
title: Update Workers VPC connectivity service
page_id: operation-put-accounts-account-id-connectivity-directory-services-service-id-de10f38a
path: operations/connectivity-services
description: Update Workers VPC connectivity service
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/connectivity/directory/services/{service_id}
operation_ids:
    - connectivity-services-put
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Workers VPC connectivity service

`PUT /accounts/{account_id}/connectivity/directory/services/{service_id}`

Operation ID: `connectivity-services-put`

## Definition

```yaml
{"operationId": "connectivity-services-put", "summary": "Update Workers VPC connectivity service", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "service_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"tcp-with-port": {"summary": "TCP service with explicit port", "value": {"host": {"ipv4": "10.0.0.1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "postgres-db", "tcp_port": 5432, "type": "tcp"}}, "tcp-with-protocol": {"summary": "TCP service with app protocol", "value": {"app_protocol": "postgresql", "host": {"hostname": "db.example.com", "resolver_network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "postgres-db", "tcp_port": 5432, "type": "tcp"}}}, "schema": {"$ref": "#/components/schemas/infra_ServiceConfig"}}}}, "responses": {"200": {"description": "Successfully updated Workers VPC connectivity service", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_ServiceConfig"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to update Workers VPC connectivity service", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Connectivity Services"]}
```
