---
title: Delete Custom Origin Trust Store
page_id: operation-delete-zones-zone-id-acm-custom-trust-store-custom-origin-trust-store-id-f1e8e1bc
path: operations/custom-origin-trust-store
description: Removes a root CA certificate from the custom origin trust store. Origins using certificates signed by this CA will no longer be trusted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/acm/custom_trust_store/{custom_origin_trust_store_id}
operation_ids:
    - custom-origin-trust-store-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Custom Origin Trust Store

`DELETE /zones/{zone_id}/acm/custom_trust_store/{custom_origin_trust_store_id}`

Operation ID: `custom-origin-trust-store-delete`

Removes a root CA certificate from the custom origin trust store. Origins using certificates signed by this CA will no longer be trusted.

## Definition

```yaml
{"operationId": "custom-origin-trust-store-delete", "summary": "Delete Custom Origin Trust Store", "description": "Removes a root CA certificate from the custom origin trust store. Origins using certificates signed by this CA will no longer be trusted.", "parameters": [{"name": "custom_origin_trust_store_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Delete Custom Origin Trust Store response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_trust_store_response_id_only"}}}}, "4XX": {"description": "Delete Custom Origin Trust Store response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Origin Trust Store"], "x-api-token-group": ["SSL and Certificates Write"], "x-cfPermissionsRequired": {"enum": ["#ssl:read", "#ssl:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "acm.custom-trust-store", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
