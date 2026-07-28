---
title: Delete Custom Hostname (and any issued SSL certificates)
page_id: operation-delete-zones-zone-id-custom-hostnames-custom-hostname-id-2d871822
path: operations/custom-hostname-for-a-zone
description: Permanently deletes a custom hostname and revokes any SSL certificates that were issued for it. This action cannot be undone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/custom_hostnames/{custom_hostname_id}
operation_ids:
    - custom-hostname-for-a-zone-delete-custom-hostname-(-and-any-issued-ssl-certificates)
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Custom Hostname (and any issued SSL certificates)

`DELETE /zones/{zone_id}/custom_hostnames/{custom_hostname_id}`

Operation ID: `custom-hostname-for-a-zone-delete-custom-hostname-(-and-any-issued-ssl-certificates)`

Permanently deletes a custom hostname and revokes any SSL certificates that were issued for it. This action cannot be undone.

## Definition

```yaml
{"operationId": "custom-hostname-for-a-zone-delete-custom-hostname-(-and-any-issued-ssl-certificates)", "summary": "Delete Custom Hostname (and any issued SSL certificates)", "description": "Permanently deletes a custom hostname and revokes any SSL certificates that were issued for it. This action cannot be undone.", "parameters": [{"name": "custom_hostname_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Custom Hostname (and any issued SSL certificates) response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}}}}}, "4XX": {"description": "Delete Custom Hostname (and any issued SSL certificates) response failure.", "content": {"application/json": {"schema": {"allOf": [{"properties": {"id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, "type": "object"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Hostname for a Zone"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-hostnames", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
