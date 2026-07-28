---
title: List Web3 Hostnames
page_id: operation-get-zones-zone-id-web3-hostnames-19dd19c7
path: operations/web3-hostname
description: List Web3 Hostnames
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/web3/hostnames
operation_ids:
    - web3-hostname-list-web3-hostnames
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Web3 Hostnames

`GET /zones/{zone_id}/web3/hostnames`

Operation ID: `web3-hostname-list-web3-hostnames`

## Definition

```yaml
{"operationId": "web3-hostname-list-web3-hostnames", "summary": "List Web3 Hostnames", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "responses": {"200": {"description": "List Web3 Hostnames response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_collection_response"}}}}, "4XX": {"description": "List Web3 Hostnames error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_collection_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "List Web3 Hostnames response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_collection_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-api-token-group": ["Web3 Hostnames Write", "Web3 Hostnames Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
