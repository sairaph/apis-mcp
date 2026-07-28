---
title: Update multiple GRE tunnels
page_id: operation-put-accounts-account-id-magic-gre-tunnels-93d8e327
path: operations/magic-gre-tunnels
description: Updates multiple GRE tunnels. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/gre_tunnels
operation_ids:
    - magic-gre-tunnels-update-multiple-gre-tunnels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update multiple GRE tunnels

`PUT /accounts/{account_id}/magic/gre_tunnels`

Operation ID: `magic-gre-tunnels-update-multiple-gre-tunnels`

Updates multiple GRE tunnels. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.

## Definition

```yaml
{"operationId": "magic-gre-tunnels-update-multiple-gre-tunnels", "summary": "Update multiple GRE tunnels", "description": "Updates multiple GRE tunnels. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the request and response bodies will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"required": ["id"]}}}}, "responses": {"200": {"description": "Update multiple GRE tunnels response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_modified_tunnels_collection_response"}}}}, "4XX": {"description": "Update multiple GRE tunnels response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_modified_tunnels_collection_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic GRE tunnels"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.gre-tunnels", "x-fern-sdk-method-name": "bulk-update", "x-forge-hidden": true}
```
