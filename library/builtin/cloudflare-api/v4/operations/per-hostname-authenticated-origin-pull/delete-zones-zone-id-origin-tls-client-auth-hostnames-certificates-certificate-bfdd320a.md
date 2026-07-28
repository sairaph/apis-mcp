---
title: Delete Hostname Client Certificate
page_id: operation-delete-zones-zone-id-origin-tls-client-auth-hostnames-certificates-certi-a67ef1de
path: operations/per-hostname-authenticated-origin-pull
description: |-
    Removes a client certificate used for authenticated origin pulls on a specific hostname.
    Note: Before deleting the certificate, you must first invalidate the hostname for client authentication by sending a PUT request with `enabled` set to null. After invalidating the association, the certificate can be safely deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates/{certificate_id}
operation_ids:
    - per-hostname-authenticated-origin-pull-delete-hostname-client-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Hostname Client Certificate

`DELETE /zones/{zone_id}/origin_tls_client_auth/hostnames/certificates/{certificate_id}`

Operation ID: `per-hostname-authenticated-origin-pull-delete-hostname-client-certificate`

Removes a client certificate used for authenticated origin pulls on a specific hostname.
Note: Before deleting the certificate, you must first invalidate the hostname for client authentication by sending a PUT request with `enabled` set to null. After invalidating the association, the certificate can be safely deleted.

## Definition

```yaml
{"operationId": "per-hostname-authenticated-origin-pull-delete-hostname-client-certificate", "summary": "Delete Hostname Client Certificate", "description": "Removes a client certificate used for authenticated origin pulls on a specific hostname.\nNote: Before deleting the certificate, you must first invalidate the hostname for client authentication by sending a PUT request with `enabled` set to null. After invalidating the association, the certificate can be safely deleted.\n", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Hostname Client Certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-4"}}}}, "4XX": {"description": "Delete Hostname Client Certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-4"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Per-hostname Authenticated Origin Pull"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.hostname-certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
