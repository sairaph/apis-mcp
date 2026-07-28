---
title: Update IPFS Universal Path Gateway Content List
page_id: operation-put-zones-zone-id-web3-hostnames-identifier-ipfs-universal-path-content-03a889d6
path: operations/web3-hostname
description: Update IPFS Universal Path Gateway Content List
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list
operation_ids:
    - web3-hostname-update-ipfs-universal-path-gateway-content-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update IPFS Universal Path Gateway Content List

`PUT /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list`

Operation ID: `web3-hostname-update-ipfs-universal-path-gateway-content-list`

## Definition

```yaml
{"operationId": "web3-hostname-update-ipfs-universal-path-gateway-content-list", "summary": "Update IPFS Universal Path Gateway Content List", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_content_list_update_request"}}}}, "responses": {"200": {"description": "Update IPFS Universal Path Gateway Content List response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_content_list_details_response"}}}}, "4XX": {"description": "Update IPFS Universal Path Gateway Content List error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_details_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "Update IPFS Universal Path Gateway Content List response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_details_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-api-token-group": ["Web3 Hostnames Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
