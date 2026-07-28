---
title: SSL Verification Details
page_id: operation-get-zones-zone-id-ssl-verification-45e542cc
path: operations/ssl-verification
description: Get SSL Verification Info for a Zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ssl/verification
operation_ids:
    - ssl-verification-ssl-verification-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# SSL Verification Details

`GET /zones/{zone_id}/ssl/verification`

Operation ID: `ssl-verification-ssl-verification-details`

Get SSL Verification Info for a Zone.

## Definition

```yaml
{"operationId": "ssl-verification-ssl-verification-details", "summary": "SSL Verification Details", "description": "Get SSL Verification Info for a Zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "retry", "in": "query", "schema": {"description": "Immediately retry SSL Verification.", "type": "boolean", "example": true, "enum": [true]}}], "responses": {"200": {"description": "SSL Verification Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_verification_response_collection"}}}}, "4XX": {"description": "SSL Verification Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_ssl_verification_response_collection"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["SSL Verification"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read", "SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ssl.verification", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
