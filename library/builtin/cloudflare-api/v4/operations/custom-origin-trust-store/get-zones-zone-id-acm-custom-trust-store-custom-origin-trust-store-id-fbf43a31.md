---
title: Custom Origin Trust Store Details
page_id: operation-get-zones-zone-id-acm-custom-trust-store-custom-origin-trust-store-id-86766d42
path: operations/custom-origin-trust-store
description: Retrieves details about a specific root CA certificate in the custom origin trust store, including expiration and subject information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/acm/custom_trust_store/{custom_origin_trust_store_id}
operation_ids:
    - custom-origin-trust-store-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Custom Origin Trust Store Details

`GET /zones/{zone_id}/acm/custom_trust_store/{custom_origin_trust_store_id}`

Operation ID: `custom-origin-trust-store-details`

Retrieves details about a specific root CA certificate in the custom origin trust store, including expiration and subject information.

## Definition

```yaml
{"operationId": "custom-origin-trust-store-details", "summary": "Custom Origin Trust Store Details", "description": "Retrieves details about a specific root CA certificate in the custom origin trust store, including expiration and subject information.", "parameters": [{"name": "custom_origin_trust_store_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_identifier"}}], "responses": {"200": {"description": "Custom Origin Trust Store Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_trust_store_response_single"}}}}, "4XX": {"description": "Custom Origin Trust Store Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_trust_store_response_single"}, {"$ref": "#/components/schemas/tls-certificates-and-hostnames_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Custom Origin Trust Store"], "x-api-token-group": ["SSL and Certificates Write", "SSL and Certificates Read"], "x-cfPermissionsRequired": {"enum": ["#ssl:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "acm.custom-trust-store", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
