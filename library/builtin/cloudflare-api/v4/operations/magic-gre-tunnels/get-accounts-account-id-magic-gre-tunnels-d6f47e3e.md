---
title: List GRE tunnels
page_id: operation-get-accounts-account-id-magic-gre-tunnels-4a04bd85
path: operations/magic-gre-tunnels
description: Lists GRE tunnels associated with an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/gre_tunnels
operation_ids:
    - magic-gre-tunnels-list-gre-tunnels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List GRE tunnels

`GET /accounts/{account_id}/magic/gre_tunnels`

Operation ID: `magic-gre-tunnels-list-gre-tunnels`

Lists GRE tunnels associated with an account.

## Definition

```yaml
{"operationId": "magic-gre-tunnels-list-gre-tunnels", "summary": "List GRE tunnels", "description": "Lists GRE tunnels associated with an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the response body will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "List GRE tunnels response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_tunnels_collection_response"}}}}, "4XX": {"description": "List GRE tunnels response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_tunnels_collection_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic GRE tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.gre-tunnels", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
