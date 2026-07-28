---
title: Edit Custom Hostname
page_id: operation-patch-zones-zone-id-custom-hostnames-custom-hostname-id-652fe3ae
path: operations/custom-hostname-for-a-zone
description: Modify SSL configuration for a custom hostname. When sent with SSL config that matches existing config, used to indicate that hostname should pass domain control validation (DCV). Can also be used to change validation type, e.g., from 'http' to 'email'. Bundle an existing certificate with another certificate by using the "custom_cert_bundle" field. The bundling process supports combining certificates as long as the following condition is met. One certificate must use the RSA algorithm, and the other must use the ECDSA algorithm.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/{custom_hostname_id}
operation_ids:
    - custom-hostname-for-a-zone-edit-custom-hostname
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Custom Hostname

`PATCH /zones/{zone_id}/custom_hostnames/{custom_hostname_id}`

Operation ID: `custom-hostname-for-a-zone-edit-custom-hostname`

Modify SSL configuration for a custom hostname. When sent with SSL config that matches existing config, used to indicate that hostname should pass domain control validation (DCV). Can also be used to change validation type, e.g., from 'http' to 'email'. Bundle an existing certificate with another certificate by using the "custom_cert_bundle" field. The bundling process supports combining certificates as long as the following condition is met. One certificate must use the RSA algorithm, and the other must use the ECDSA algorithm.

## Definition

```yaml
{"operationId": "custom-hostname-for-a-zone-edit-custom-hostname", "summary": "Edit Custom Hostname", "description": "Modify SSL configuration for a custom hostname. When sent with SSL config that matches existing config, used to indicate that hostname should pass domain control validation (DCV). Can also be used to change validation type, e.g., from 'http' to 'email'. Bundle an existing certificate with another certificate by using the \"custom_cert_bundle\" field. The bundling process supports combining certificates as long as the following condition is met. One certificate must use the RSA algorithm, and the other must use the ECDSA algorithm.", "parameters": [{"name": "custom_hostname_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"custom_metadata": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_metadata"}, "custom_origin_server": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_origin_server"}, "custom_origin_sni": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_origin_sni"}, "ssl": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_sslpost"}}}}}}, "responses": {"200": {"description": "Edit Custom Hostname response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}}}}, "4XX": {"description": "Edit Custom Hostname response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
