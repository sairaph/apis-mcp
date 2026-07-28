---
title: Get hostname route
page_id: operation-get-accounts-account-id-zerotrust-routes-hostname-hostname-route-id-e2f91d20
path: operations/zero-trust-hostname-route
description: Get a hostname route.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zerotrust/routes/hostname/{hostname_route_id}
operation_ids:
    - zero-trust-networks-route-hostname-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get hostname route

`GET /accounts/{account_id}/zerotrust/routes/hostname/{hostname_route_id}`

Operation ID: `zero-trust-networks-route-hostname-get`

Get a hostname route.

## Definition

```yaml
{"operationId": "zero-trust-networks-route-hostname-get", "summary": "Get hostname route", "description": "Get a hostname route.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "hostname_route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_hostname_route_id"}}], "responses": {"200": {"description": "Get hostname route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_hostname_route_response_single"}}}}, "4XX": {"description": "Get hostname route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_hostname_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Hostname Route"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.hostname-routes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
