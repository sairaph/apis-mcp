---
title: Create Client Certificate
page_id: operation-post-zones-zone-id-client-certificates-0b571d4c
path: operations/api-shield-client-certificates-for-a-zone
description: Create a new API Shield mTLS Client Certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/client_certificates
operation_ids:
    - client-certificate-for-a-zone-create-client-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Client Certificate

`POST /zones/{zone_id}/client_certificates`

Operation ID: `client-certificate-for-a-zone-create-client-certificate`

Create a new API Shield mTLS Client Certificate.

## Definition

```yaml
{"operationId": "client-certificate-for-a-zone-create-client-certificate", "summary": "Create Client Certificate", "description": "Create a new API Shield mTLS Client Certificate.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"csr": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_csr-2"}, "validity_days": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_validity_days-2"}}, "required": ["csr", "validity_days"]}}}}, "responses": {"200": {"description": "Create Client Certificate Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_client_certificate_response_single"}}}}, "4XX": {"description": "Create Client Certificate Response Failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_client_certificate_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Client Certificates for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "client-certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
