---
title: Create hostname route
page_id: operation-post-accounts-account-id-zerotrust-routes-hostname-bfd6ff68
path: operations/zero-trust-hostname-route
description: Create a hostname route.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/zerotrust/routes/hostname
operation_ids:
    - zero-trust-networks-route-hostname-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create hostname route

`POST /accounts/{account_id}/zerotrust/routes/hostname`

Operation ID: `zero-trust-networks-route-hostname-create`

Create a hostname route.

## Definition

```yaml
{"operationId": "zero-trust-networks-route-hostname-create", "summary": "Create hostname route", "description": "Create a hostname route.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_hostname_comment"}, "hostname": {"$ref": "#/components/schemas/tunnel_hostname"}, "tunnel_id": {"$ref": "#/components/schemas/tunnel_tunnel_id-3"}}}}}}, "responses": {"200": {"description": "Create hostname route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_hostname_route_response_single"}}}}, "4XX": {"description": "Create hostname route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_hostname_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Hostname Route"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.hostname-routes", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
