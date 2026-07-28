---
title: Upload Custom Origin Trust Store
page_id: operation-post-zones-zone-id-acm-custom-trust-store-8bbf5ece
path: operations/custom-origin-trust-store
description: Upload a root CA certificate to the Custom Origin Trust Store for a Zone. Only root CA certificates are accepted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/acm/custom_trust_store
operation_ids:
    - custom-origin-trust-store-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload Custom Origin Trust Store

`POST /zones/{zone_id}/acm/custom_trust_store`

Operation ID: `custom-origin-trust-store-create`

Upload a root CA certificate to the Custom Origin Trust Store for a Zone. Only root CA certificates are accepted.

## Definition

```yaml
{"operationId": "custom-origin-trust-store-create", "summary": "Upload Custom Origin Trust Store", "description": "Upload a root CA certificate to the Custom Origin Trust Store for a Zone. Only root CA certificates are accepted.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate-3"}}, "required": ["certificate"]}}}}, "responses": {"200": {"description": "Upload Custom Origin Trust Store response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_trust_store_response_single"}}}}, "4XX": {"description": "Upload Custom Origin Trust Store response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_trust_store_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Origin Trust Store"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read", "#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "acm.custom-trust-store", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
