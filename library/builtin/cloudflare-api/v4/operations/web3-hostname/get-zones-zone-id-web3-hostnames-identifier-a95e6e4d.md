---
title: Web3 Hostname Details
page_id: operation-get-zones-zone-id-web3-hostnames-identifier-05618686
path: operations/web3-hostname
description: Web3 Hostname Details
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/web3/hostnames/{identifier}
operation_ids:
    - web3-hostname-web3-hostname-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Web3 Hostname Details

`GET /zones/{zone_id}/web3/hostnames/{identifier}`

Operation ID: `web3-hostname-web3-hostname-details`

## Definition

```yaml
{"operationId": "web3-hostname-web3-hostname-details", "summary": "Web3 Hostname Details", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "responses": {"200": {"description": "Web3 Hostname Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_single_response"}}}}, "4XX": {"description": "Web3 Hostname Details error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_collection_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "Web3 Hostname Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_single_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-api-token-group": ["Web3 Hostnames Write", "Web3 Hostnames Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
