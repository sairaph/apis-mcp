---
title: Revoke Client Certificate
page_id: operation-delete-zones-zone-id-client-certificates-client-certificate-id-a64a769a
path: operations/api-shield-client-certificates-for-a-zone
description: Set a API Shield mTLS Client Certificate to pending_revocation status for processing to revoked status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/client_certificates/{client_certificate_id}
operation_ids:
    - client-certificate-for-a-zone-delete-client-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke Client Certificate

`DELETE /zones/{zone_id}/client_certificates/{client_certificate_id}`

Operation ID: `client-certificate-for-a-zone-delete-client-certificate`

Set a API Shield mTLS Client Certificate to pending_revocation status for processing to revoked status.

## Definition

```yaml
{"operationId": "client-certificate-for-a-zone-delete-client-certificate", "summary": "Revoke Client Certificate", "description": "Set a API Shield mTLS Client Certificate to pending_revocation status for processing to revoked status.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "client_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Revoke Client Certificate Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_client_certificate_response_single"}}}}, "4XX": {"description": "Revoke Client Certificate Response Failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Client Certificates for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "client-certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
