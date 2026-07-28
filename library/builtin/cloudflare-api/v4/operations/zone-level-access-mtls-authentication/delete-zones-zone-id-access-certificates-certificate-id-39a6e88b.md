---
title: Delete an mTLS certificate
page_id: operation-delete-zones-zone-id-access-certificates-certificate-id-64abf307
path: operations/zone-level-access-mtls-authentication
description: Deletes an mTLS certificate.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/access/certificates/{certificate_id}
operation_ids:
    - zone-level-access-mtls-authentication-delete-an-mtls-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an mTLS certificate

`DELETE /zones/{zone_id}/access/certificates/{certificate_id}`

Operation ID: `zone-level-access-mtls-authentication-delete-an-mtls-certificate`

Deletes an mTLS certificate.

## Definition

```yaml
{"operationId": "zone-level-access-mtls-authentication-delete-an-mtls-certificate", "summary": "Delete an mTLS certificate", "description": "Deletes an mTLS certificate.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Delete an mTLS certificate response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response-3"}}}}, "4XX": {"description": "Delete an mTLS certificate response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access mTLS authentication"], "x-api-token-group": ["Access: Mutual TLS Certificates Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
