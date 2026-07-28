---
title: IPFS Universal Path Gateway Content List Details
page_id: operation-get-zones-zone-id-web3-hostnames-identifier-ipfs-universal-path-content-38f3c7a3
path: operations/web3-hostname
description: IPFS Universal Path Gateway Content List Details
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list
operation_ids:
    - web3-hostname-ipfs-universal-path-gateway-content-list-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# IPFS Universal Path Gateway Content List Details

`GET /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list`

Operation ID: `web3-hostname-ipfs-universal-path-gateway-content-list-details`

## Definition

```yaml
{"operationId": "web3-hostname-ipfs-universal-path-gateway-content-list-details", "summary": "IPFS Universal Path Gateway Content List Details", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "responses": {"200": {"description": "IPFS Universal Path Gateway Content List Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_content_list_details_response"}}}}, "4XX": {"description": "IPFS Universal Path Gateway Content List Details error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_details_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "IPFS Universal Path Gateway Content List Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_details_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-api-token-group": ["Web3 Hostnames Write", "Web3 Hostnames Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
