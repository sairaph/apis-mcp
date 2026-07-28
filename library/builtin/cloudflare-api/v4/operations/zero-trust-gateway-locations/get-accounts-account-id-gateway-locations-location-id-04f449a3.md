---
title: Get Zero Trust Gateway location details
page_id: operation-get-accounts-account-id-gateway-locations-location-id-ec160f0e
path: operations/zero-trust-gateway-locations
description: Get a single Zero Trust Gateway location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/locations/{location_id}
operation_ids:
    - zero-trust-gateway-locations-zero-trust-gateway-location-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust Gateway location details

`GET /accounts/{account_id}/gateway/locations/{location_id}`

Operation ID: `zero-trust-gateway-locations-zero-trust-gateway-location-details`

Get a single Zero Trust Gateway location.

## Definition

```yaml
{"operationId": "zero-trust-gateway-locations-zero-trust-gateway-location-details", "summary": "Get Zero Trust Gateway location details", "description": "Get a single Zero Trust Gateway location.", "parameters": [{"name": "location_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Gets Zero Trust Gateway location details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-3"}}}}, "4XX": {"description": "Gets Zero Trust Gateway location details response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-3"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway locations"], "x-api-token-group": ["Cloudflare Zero Trust Secure DNS Locations Write", "Zero Trust Read", "Zero Trust Write"]}
```
