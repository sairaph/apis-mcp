---
title: Update interconnect
page_id: operation-put-accounts-account-id-magic-cf-interconnects-cf-interconnect-id-81b09c40
path: operations/magic-interconnects
description: Updates a specific interconnect associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/cf_interconnects/{cf_interconnect_id}
operation_ids:
    - magic-interconnects-update-interconnect
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update interconnect

`PUT /accounts/{account_id}/magic/cf_interconnects/{cf_interconnect_id}`

Operation ID: `magic-interconnects-update-interconnect`

Updates a specific interconnect associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.

## Definition

```yaml
{"operationId": "magic-interconnects-update-interconnect", "summary": "Update interconnect", "description": "Updates a specific interconnect associated with an account. Use `?validate_only=true` as an optional query parameter to only run validation without persisting changes.", "parameters": [{"name": "cf_interconnect_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "x-magic-new-hc-target", "in": "header", "description": "If true, the health check target in the request and response bodies will be presented using the new object format. Defaults to false.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_interconnect_tunnel_update_request"}}}}, "responses": {"200": {"description": "Update interconnect response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_components-schemas-tunnel_modified_response"}}}}, "4XX": {"description": "Update interconnect response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_components-schemas-tunnel_modified_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Interconnects"], "x-api-token-group": ["Magic WAN Write", "Magic Transit Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "magic-transit.cf-interconnects", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
