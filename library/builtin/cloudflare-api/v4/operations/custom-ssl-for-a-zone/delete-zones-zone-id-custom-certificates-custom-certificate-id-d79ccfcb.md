---
title: Delete SSL Configuration
page_id: operation-delete-zones-zone-id-custom-certificates-custom-certificate-id-e75b848a
path: operations/custom-ssl-for-a-zone
description: Remove a SSL certificate from a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/custom_certificates/{custom_certificate_id}
operation_ids:
    - custom-ssl-for-a-zone-delete-ssl-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete SSL Configuration

`DELETE /zones/{zone_id}/custom_certificates/{custom_certificate_id}`

Operation ID: `custom-ssl-for-a-zone-delete-ssl-configuration`

Remove a SSL certificate from a zone.

## Definition

```yaml
{"operationId": "custom-ssl-for-a-zone-delete-ssl-configuration", "summary": "Delete SSL Configuration", "description": "Remove a SSL certificate from a zone.", "parameters": [{"name": "custom_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete SSL Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_id_only"}}}}, "4XX": {"description": "Delete SSL Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_id_only"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom SSL for a Zone"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-certificates", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
