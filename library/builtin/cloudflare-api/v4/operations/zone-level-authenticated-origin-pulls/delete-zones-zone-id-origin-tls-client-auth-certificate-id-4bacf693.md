---
title: Delete Certificate
page_id: operation-delete-zones-zone-id-origin-tls-client-auth-certificate-id-c649a3af
path: operations/zone-level-authenticated-origin-pulls
description: Removes a client certificate used for zone-level authenticated origin pulls.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/origin_tls_client_auth/{certificate_id}
operation_ids:
    - zone-level-authenticated-origin-pulls-delete-certificate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Certificate

`DELETE /zones/{zone_id}/origin_tls_client_auth/{certificate_id}`

Operation ID: `zone-level-authenticated-origin-pulls-delete-certificate`

Removes a client certificate used for zone-level authenticated origin pulls.

## Definition

```yaml
{"operationId": "zone-level-authenticated-origin-pulls-delete-certificate", "summary": "Delete Certificate", "description": "Removes a client certificate used for zone-level authenticated origin pulls.", "parameters": [{"name": "certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Certificate response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-3"}}}}, "4XX": {"description": "Delete Certificate response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single-3"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone-Level Authenticated Origin Pulls"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "origin-tls-client-auth.zone-certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
