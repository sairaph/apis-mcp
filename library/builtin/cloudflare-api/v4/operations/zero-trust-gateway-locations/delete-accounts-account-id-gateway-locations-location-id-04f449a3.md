---
title: Delete a Zero Trust Gateway location
page_id: operation-delete-accounts-account-id-gateway-locations-location-id-3c4a2394
path: operations/zero-trust-gateway-locations
description: Delete a configured Zero Trust Gateway location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/gateway/locations/{location_id}
operation_ids:
    - zero-trust-gateway-locations-delete-zero-trust-gateway-location
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Zero Trust Gateway location

`DELETE /accounts/{account_id}/gateway/locations/{location_id}`

Operation ID: `zero-trust-gateway-locations-delete-zero-trust-gateway-location`

Delete a configured Zero Trust Gateway location.

## Definition

```yaml
{"operationId": "zero-trust-gateway-locations-delete-zero-trust-gateway-location", "summary": "Delete a Zero Trust Gateway location", "description": "Delete a configured Zero Trust Gateway location.", "parameters": [{"name": "location_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Deletes a Zero Trust Gateway location response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}}}}, "4XX": {"description": "Deletes a Zero Trust Gateway location response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway locations"], "x-api-token-group": ["Cloudflare Zero Trust Secure DNS Locations Write", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.locations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
