---
title: Reactivate Client Certificate
page_id: operation-patch-zones-zone-id-client-certificates-client-certificate-id-92e249af
path: operations/api-shield-client-certificates-for-a-zone
description: If a API Shield mTLS Client Certificate is in a pending_revocation state, you may reactivate it with this endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/client_certificates/{client_certificate_id}
operation_ids:
    - client-certificate-for-a-zone-edit-client-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reactivate Client Certificate

`PATCH /zones/{zone_id}/client_certificates/{client_certificate_id}`

Operation ID: `client-certificate-for-a-zone-edit-client-certificate`

If a API Shield mTLS Client Certificate is in a pending_revocation state, you may reactivate it with this endpoint.

## Definition

```yaml
{"operationId": "client-certificate-for-a-zone-edit-client-certificate", "summary": "Reactivate Client Certificate", "description": "If a API Shield mTLS Client Certificate is in a pending_revocation state, you may reactivate it with this endpoint.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "client_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"reactivate": {"type": "boolean", "example": true}}}}}}, "responses": {"200": {"description": "Reactivate Client Certificate Response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_client_certificate_response_single"}}}}, "4XX": {"description": "Reactivate Client Certificate Response Failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["API Shield Client Certificates for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "client-certificates", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
