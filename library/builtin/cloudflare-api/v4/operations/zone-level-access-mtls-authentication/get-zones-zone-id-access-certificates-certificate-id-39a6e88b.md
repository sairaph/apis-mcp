---
title: Get an mTLS certificate
page_id: operation-get-zones-zone-id-access-certificates-certificate-id-8ca7f3bc
path: operations/zone-level-access-mtls-authentication
description: Fetches a single mTLS certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/certificates/{certificate_id}
operation_ids:
    - zone-level-access-mtls-authentication-get-an-mtls-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an mTLS certificate

`GET /zones/{zone_id}/access/certificates/{certificate_id}`

Operation ID: `zone-level-access-mtls-authentication-get-an-mtls-certificate`

Fetches a single mTLS certificate.

## Definition

```yaml
{"operationId": "zone-level-access-mtls-authentication-get-an-mtls-certificate", "summary": "Get an mTLS certificate", "description": "Fetches a single mTLS certificate.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an mTLS certificate response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-20"}}}}, "4XX": {"description": "Get an mTLS certificate response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.certificates", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
