---
title: Create Web3 Hostname
page_id: operation-post-zones-zone-id-web3-hostnames-55d4e131
path: operations/web3-hostname
description: Create Web3 Hostname
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/web3/hostnames
operation_ids:
    - web3-hostname-create-web3-hostname
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Web3 Hostname

`POST /zones/{zone_id}/web3/hostnames`

Operation ID: `web3-hostname-create-web3-hostname`

## Definition

```yaml
{"operationId": "web3-hostname-create-web3-hostname", "summary": "Create Web3 Hostname", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_create_request"}}}}, "responses": {"200": {"description": "Create Web3 Hostname response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_single_response"}}}}, "4XX": {"description": "Create Web3 Hostname error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_collection_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "Create Web3 Hostname response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_single_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-api-token-group": ["Web3 Hostnames Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
