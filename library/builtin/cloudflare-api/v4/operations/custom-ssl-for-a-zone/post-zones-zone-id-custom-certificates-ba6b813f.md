---
title: Create SSL Configuration
page_id: operation-post-zones-zone-id-custom-certificates-c493d4f7
path: operations/custom-ssl-for-a-zone
description: Upload a new SSL certificate for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/custom_certificates
operation_ids:
    - custom-ssl-for-a-zone-create-ssl-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create SSL Configuration

`POST /zones/{zone_id}/custom_certificates`

Operation ID: `custom-ssl-for-a-zone-create-ssl-configuration`

Upload a new SSL certificate for a zone.

## Definition

```yaml
{"operationId": "custom-ssl-for-a-zone-create-ssl-configuration", "summary": "Create SSL Configuration", "description": "Upload a new SSL certificate for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"bundle_method": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_bundle_method"}, "certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate"}, "custom_csr_id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_csr_id"}, "deploy": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_deploy"}, "geo_restrictions": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_geo_restrictions"}, "policy": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_policy"}, "private_key": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_private_key"}, "type": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_type"}}, "required": ["certificate"]}}}}, "responses": {"200": {"description": "Create SSL Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single"}}}}, "4XX": {"description": "Create SSL Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom SSL for a Zone"], "x-api-token-group": ["Access: Mutual TLS Certificates Write", "SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
