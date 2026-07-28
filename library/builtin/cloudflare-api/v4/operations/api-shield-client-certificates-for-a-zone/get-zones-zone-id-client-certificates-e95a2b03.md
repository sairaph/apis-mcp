---
title: List Client Certificates
page_id: operation-get-zones-zone-id-client-certificates-6591dfdc
path: operations/api-shield-client-certificates-for-a-zone
description: List all of your Zone's API Shield mTLS Client Certificates by Status and/or using Pagination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/client_certificates
operation_ids:
    - client-certificate-for-a-zone-list-client-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Client Certificates

`GET /zones/{zone_id}/client_certificates`

Operation ID: `client-certificate-for-a-zone-list-client-certificates`

List all of your Zone's API Shield mTLS Client Certificates by Status and/or using Pagination.

## Definition

```yaml
{"operationId": "client-certificate-for-a-zone-list-client-certificates", "summary": "List Client Certificates", "description": "List all of your Zone's API Shield mTLS Client Certificates by Status and/or using Pagination.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "status", "in": "query", "schema": {"description": "Client Certitifcate Status to filter results by.", "type": "string", "example": "all", "enum": ["all", "active", "pending_reactivation", "pending_revocation", "revoked"]}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of records per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "limit", "in": "query", "schema": {"description": "Limit to the number of records returned.", "type": "integer", "example": 10}}, {"name": "offset", "in": "query", "schema": {"description": "Offset the results.", "type": "integer", "example": 10}}], "responses": {"200": {"description": "List Client Certificates Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_client_certificate_response_collection"}}}}, "4XX": {"description": "List Client Certificates Response Failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Client Certificates for a Zone"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "client-certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
