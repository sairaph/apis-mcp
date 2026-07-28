---
title: List hostname routes
page_id: operation-get-accounts-account-id-zerotrust-routes-hostname-d8cb9a6a
path: operations/zero-trust-hostname-route
description: Lists and filters hostname routes in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zerotrust/routes/hostname
operation_ids:
    - zero-trust-networks-route-hostname-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List hostname routes

`GET /accounts/{account_id}/zerotrust/routes/hostname`

Operation ID: `zero-trust-networks-route-hostname-list`

Lists and filters hostname routes in an account.

## Definition

```yaml
{"operationId": "zero-trust-networks-route-hostname-list", "summary": "List hostname routes", "description": "Lists and filters hostname routes in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_hostname_route_id"}}, {"name": "hostname", "in": "query", "description": "If set, only list hostname routes that contain a substring of the given value, the filter is case-insensitive.", "schema": {"$ref": "#/components/schemas/tunnel_hostname"}}, {"name": "tunnel_id", "in": "query", "description": "If set, only list hostname routes that point to a specific tunnel.", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id-3"}}, {"name": "comment", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_hostname_query_comment"}}, {"name": "existed_at", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_existed_at"}}, {"name": "is_deleted", "in": "query", "schema": {"description": "If `true`, only return deleted hostname routes. If `false`, exclude deleted hostname routes.", "type": "boolean", "example": true, "default": false}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_per_page"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_page_number"}}], "responses": {"200": {"description": "List hostname routes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_hostname_route_response_collection"}}}}, "4XX": {"description": "List hostname routes failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_hostname_route_response_collection"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Hostname Route"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.hostname-routes", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
