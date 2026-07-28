---
title: Edit Keyless SSL Configuration
page_id: operation-patch-zones-zone-id-keyless-certificates-keyless-certificate-id-e4bdbd7e
path: operations/keyless-ssl-for-a-zone
description: 'This will update attributes of a Keyless SSL. Consists of one or more of the following:  host,name,port.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/keyless_certificates/{keyless_certificate_id}
operation_ids:
    - keyless-ssl-for-a-zone-edit-keyless-ssl-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Keyless SSL Configuration

`PATCH /zones/{zone_id}/keyless_certificates/{keyless_certificate_id}`

Operation ID: `keyless-ssl-for-a-zone-edit-keyless-ssl-configuration`

This will update attributes of a Keyless SSL. Consists of one or more of the following:  host,name,port.

## Definition

```yaml
{"operationId": "keyless-ssl-for-a-zone-edit-keyless-ssl-configuration", "summary": "Edit Keyless SSL Configuration", "description": "This will update attributes of a Keyless SSL. Consists of one or more of the following:  host,name,port.", "parameters": [{"name": "keyless_certificate_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"enabled": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_enabled_write"}, "host": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_host"}, "name": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_name_write"}, "port": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_port"}, "tunnel": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_tunnel"}}}}}}, "responses": {"200": {"description": "Edit Keyless SSL Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single"}}}}, "4XX": {"description": "Edit Keyless SSL Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Keyless SSL for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "keyless-certificates", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
