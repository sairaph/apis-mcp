---
title: Create a Zero Trust Gateway location
page_id: operation-post-accounts-account-id-gateway-locations-38b0724b
path: operations/zero-trust-gateway-locations
description: Create a new Zero Trust Gateway location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/locations
operation_ids:
    - zero-trust-gateway-locations-create-zero-trust-gateway-location
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Zero Trust Gateway location

`POST /accounts/{account_id}/gateway/locations`

Operation ID: `zero-trust-gateway-locations-create-zero-trust-gateway-location`

Create a new Zero Trust Gateway location.

## Definition

```yaml
{"operationId": "zero-trust-gateway-locations-create-zero-trust-gateway-location", "summary": "Create a Zero Trust Gateway location", "description": "Create a new Zero Trust Gateway location.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"client_default": {"$ref": "#/components/schemas/zero-trust-gateway_client-default"}, "dns_destination_ips_id": {"$ref": "#/components/schemas/zero-trust-gateway_dns-destination-ips-id-write"}, "ecs_support": {"$ref": "#/components/schemas/zero-trust-gateway_ecs-support"}, "endpoints": {"$ref": "#/components/schemas/zero-trust-gateway_endpoints"}, "max_ttl": {"$ref": "#/components/schemas/zero-trust-gateway_max-ttl"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-2"}, "networks": {"$ref": "#/components/schemas/zero-trust-gateway_ipv4_networks"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Creates a Zero Trust Gateway location response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-3"}}}}, "4XX": {"description": "Creates a Zero Trust Gateway location response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-3"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway locations"], "x-api-token-group": ["Cloudflare Zero Trust Secure DNS Locations Write", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.locations", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
