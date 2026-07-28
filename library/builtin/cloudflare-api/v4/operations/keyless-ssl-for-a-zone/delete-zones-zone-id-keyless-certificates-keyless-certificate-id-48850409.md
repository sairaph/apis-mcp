---
title: Delete Keyless SSL Configuration
page_id: operation-delete-zones-zone-id-keyless-certificates-keyless-certificate-id-04e86b4d
path: operations/keyless-ssl-for-a-zone
description: Removes a Keyless SSL configuration. SSL connections will no longer use the keyless server for cryptographic operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/keyless_certificates/{keyless_certificate_id}
operation_ids:
    - keyless-ssl-for-a-zone-delete-keyless-ssl-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Keyless SSL Configuration

`DELETE /zones/{zone_id}/keyless_certificates/{keyless_certificate_id}`

Operation ID: `keyless-ssl-for-a-zone-delete-keyless-ssl-configuration`

Removes a Keyless SSL configuration. SSL connections will no longer use the keyless server for cryptographic operations.

## Definition

```yaml
{"operationId": "keyless-ssl-for-a-zone-delete-keyless-ssl-configuration", "summary": "Delete Keyless SSL Configuration", "description": "Removes a Keyless SSL configuration. SSL connections will no longer use the keyless server for cryptographic operations.", "parameters": [{"name": "keyless_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Keyless SSL Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single_id"}}}}, "4XX": {"description": "Delete Keyless SSL Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single_id"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Keyless SSL for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keyless-certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
