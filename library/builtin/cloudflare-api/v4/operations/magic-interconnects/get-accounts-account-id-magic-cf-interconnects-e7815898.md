---
title: List interconnects
page_id: operation-get-accounts-account-id-magic-cf-interconnects-f33badcf
path: operations/magic-interconnects
description: Lists interconnects associated with an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/cf_interconnects
operation_ids:
    - magic-interconnects-list-interconnects
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List interconnects

`GET /accounts/{account_id}/magic/cf_interconnects`

Operation ID: `magic-interconnects-list-interconnects`

Lists interconnects associated with an account.

## Definition

```yaml
{"operationId": "magic-interconnects-list-interconnects", "summary": "List interconnects", "description": "Lists interconnects associated with an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the response body will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "responses": {"200": {"description": "List interconnects response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_components-schemas-tunnels_collection_response"}}}}, "4XX": {"description": "List interconnects response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_components-schemas-tunnels_collection_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Interconnects"], "x-api-token-group": ["Magic WAN Write", "Magic WAN Read", "Magic Transit Read", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.cf-interconnects", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
