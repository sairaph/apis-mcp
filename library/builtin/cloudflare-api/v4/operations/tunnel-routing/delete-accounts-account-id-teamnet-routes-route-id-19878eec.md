---
title: Delete a tunnel route
page_id: operation-delete-accounts-account-id-teamnet-routes-route-id-4c9ace8e
path: operations/tunnel-routing
description: Deletes a private network route from an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/teamnet/routes/{route_id}
operation_ids:
    - tunnel-route-delete-a-tunnel-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a tunnel route

`DELETE /accounts/{account_id}/teamnet/routes/{route_id}`

Operation ID: `tunnel-route-delete-a-tunnel-route`

Deletes a private network route from an account.

## Definition

```yaml
{"operationId": "tunnel-route-delete-a-tunnel-route", "summary": "Delete a tunnel route", "description": "Deletes a private network route from an account.\n", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_route_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "responses": {"200": {"description": "Delete a tunnel route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_route_response_single"}}}}, "4XX": {"description": "Delete a tunnel route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.routes", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
