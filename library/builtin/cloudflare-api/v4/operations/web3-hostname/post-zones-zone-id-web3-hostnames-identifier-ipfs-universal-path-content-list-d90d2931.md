---
title: Create IPFS Universal Path Gateway Content List Entry
page_id: operation-post-zones-zone-id-web3-hostnames-identifier-ipfs-universal-path-content-69daa266
path: operations/web3-hostname
description: Create IPFS Universal Path Gateway Content List Entry
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries
operation_ids:
    - web3-hostname-create-ipfs-universal-path-gateway-content-list-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create IPFS Universal Path Gateway Content List Entry

`POST /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries`

Operation ID: `web3-hostname-create-ipfs-universal-path-gateway-content-list-entry`

## Definition

```yaml
{"operationId": "web3-hostname-create-ipfs-universal-path-gateway-content-list-entry", "summary": "Create IPFS Universal Path Gateway Content List Entry", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_content_list_entry_create_request"}}}}, "responses": {"200": {"description": "Create IPFS Universal Path Gateway Content List Entry response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_content_list_entry_single_response"}}}}, "4XX": {"description": "Create IPFS Universal Path Gateway Content List Entry error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_entry_single_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "Create IPFS Universal Path Gateway Content List Entry response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_entry_single_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
