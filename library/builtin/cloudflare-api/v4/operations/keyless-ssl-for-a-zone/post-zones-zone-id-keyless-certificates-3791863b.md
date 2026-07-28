---
title: Create Keyless SSL Configuration
page_id: operation-post-zones-zone-id-keyless-certificates-02930f4e
path: operations/keyless-ssl-for-a-zone
description: Creates a Keyless SSL configuration that allows SSL/TLS termination without exposing private keys to Cloudflare. Keys remain on your infrastructure.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/keyless_certificates
operation_ids:
    - keyless-ssl-for-a-zone-create-keyless-ssl-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Keyless SSL Configuration

`POST /zones/{zone_id}/keyless_certificates`

Operation ID: `keyless-ssl-for-a-zone-create-keyless-ssl-configuration`

Creates a Keyless SSL configuration that allows SSL/TLS termination without exposing private keys to Cloudflare. Keys remain on your infrastructure.

## Definition

```yaml
{"operationId": "keyless-ssl-for-a-zone-create-keyless-ssl-configuration", "summary": "Create Keyless SSL Configuration", "description": "Creates a Keyless SSL configuration that allows SSL/TLS termination without exposing private keys to Cloudflare. Keys remain on your infrastructure.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"bundle_method": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_bundle_method"}, "certificate": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_certificate-2"}, "host": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_host"}, "name": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_name_write"}, "port": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_port"}, "tunnel": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_tunnel"}}, "required": ["host", "port", "certificate"]}}}}, "responses": {"200": {"description": "Create Keyless SSL Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single"}}}}, "4XX": {"description": "Create Keyless SSL Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Keyless SSL for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keyless-certificates", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
