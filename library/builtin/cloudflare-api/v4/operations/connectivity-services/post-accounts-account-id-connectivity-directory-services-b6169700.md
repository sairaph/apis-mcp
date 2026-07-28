---
title: Create Workers VPC connectivity service
page_id: operation-post-accounts-account-id-connectivity-directory-services-3df83b09
path: operations/connectivity-services
description: Create Workers VPC connectivity service
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/connectivity/directory/services
operation_ids:
    - connectivity-services-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Workers VPC connectivity service

`POST /accounts/{account_id}/connectivity/directory/services`

Operation ID: `connectivity-services-post`

## Definition

```yaml
{"operationId": "connectivity-services-post", "summary": "Create Workers VPC connectivity service", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"http-with-defaults": {"summary": "HTTP service with defaults", "value": {"host": {"hostname": "api.example.com", "resolver_network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "web-server", "type": "http"}}, "http-with-hostname": {"summary": "HTTP service with hostname", "value": {"host": {"hostname": "api.example.com", "resolver_network": {"resolver_ips": ["10.0.0.1", "10.0.0.2"], "tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "https_port": 443, "name": "api-gateway", "type": "http"}}, "http-with-ipv4": {"summary": "HTTP service with IPv4 host", "value": {"host": {"ipv4": "10.0.0.1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "http_port": 8080, "https_port": 8443, "name": "web-server", "type": "http"}}, "tcp-with-port": {"summary": "TCP service with explicit port", "value": {"host": {"ipv4": "10.0.0.1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "postgres-db", "tcp_port": 5432, "type": "tcp"}}, "tcp-with-protocol": {"summary": "TCP service with app protocol", "value": {"app_protocol": "postgresql", "host": {"hostname": "db.example.com", "resolver_network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "postgres-db", "tcp_port": 5432, "type": "tcp"}}}, "schema": {"$ref": "#/components/schemas/infra_ServiceConfig"}}}}, "responses": {"200": {"description": "Successfully created Workers VPC connectivity service", "content": {"application/json": {"examples": {"http-service-created": {"summary": "Successfully created HTTP service", "value": {"errors": [], "messages": [], "result": {"created_at": "2024-01-15T09:30:00Z", "host": {"hostname": "api.example.com", "resolver_network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "web-server", "service_id": "550e8400-e29b-41d4-a716-446655440000", "type": "http", "updated_at": "2024-01-15T09:30:00Z"}, "success": true}}, "tcp-service-created": {"summary": "Successfully created TCP service", "value": {"errors": [], "messages": [], "result": {"created_at": "2024-01-15T09:30:00Z", "host": {"ipv4": "10.0.0.1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "name": "postgres-db", "service_id": "550e8400-e29b-41d4-a716-446655440001", "tcp_port": 5432, "type": "tcp", "updated_at": "2024-01-15T09:30:00Z"}, "success": true}}}, "schema": {"allOf": [{"$ref": "#/components/schemas/infra_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/infra_ServiceConfig"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to create Workers VPC connectivity service", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Connectivity Services"]}
```
