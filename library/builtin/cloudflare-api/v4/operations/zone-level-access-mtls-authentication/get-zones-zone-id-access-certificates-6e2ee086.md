---
title: List mTLS certificates
page_id: operation-get-zones-zone-id-access-certificates-3a4894b9
path: operations/zone-level-access-mtls-authentication
description: Lists all mTLS certificates.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/certificates
operation_ids:
    - zone-level-access-mtls-authentication-list-mtls-certificates
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List mTLS certificates

`GET /zones/{zone_id}/access/certificates`

Operation ID: `zone-level-access-mtls-authentication-list-mtls-certificates`

Lists all mTLS certificates.

## Definition

```yaml
{"operationId": "zone-level-access-mtls-authentication-list-mtls-certificates", "summary": "List mTLS certificates", "description": "Lists all mTLS certificates.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List mTLS certificates response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-21"}}}}, "4XX": {"description": "List mTLS certificates response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "Access: Mutual TLS Certificates Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.certificates", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
