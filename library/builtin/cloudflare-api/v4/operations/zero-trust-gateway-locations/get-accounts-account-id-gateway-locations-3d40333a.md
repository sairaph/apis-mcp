---
title: List Zero Trust Gateway locations
page_id: operation-get-accounts-account-id-gateway-locations-7f0b4776
path: operations/zero-trust-gateway-locations
description: List Zero Trust Gateway locations for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/locations
operation_ids:
    - zero-trust-gateway-locations-list-zero-trust-gateway-locations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zero Trust Gateway locations

`GET /accounts/{account_id}/gateway/locations`

Operation ID: `zero-trust-gateway-locations-list-zero-trust-gateway-locations`

List Zero Trust Gateway locations for an account.

## Definition

```yaml
{"operationId": "zero-trust-gateway-locations-list-zero-trust-gateway-locations", "summary": "List Zero Trust Gateway locations", "description": "List Zero Trust Gateway locations for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Lists Zero Trust Gateway locations response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-5"}}}}, "4XX": {"description": "Lists Zero Trust Gateway locations response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-5"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway locations"], "x-api-token-group": ["Cloudflare Zero Trust Secure DNS Locations Write", "Zero Trust Read", "Zero Trust Write"]}
```
