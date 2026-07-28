---
title: List IPFS Universal Path Gateway Content List Entries
page_id: operation-get-zones-zone-id-web3-hostnames-identifier-ipfs-universal-path-content-d127d09f
path: operations/web3-hostname
description: List IPFS Universal Path Gateway Content List Entries
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries
operation_ids:
    - web3-hostname-list-ipfs-universal-path-gateway-content-list-entries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List IPFS Universal Path Gateway Content List Entries

`GET /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries`

Operation ID: `web3-hostname-list-ipfs-universal-path-gateway-content-list-entries`

## Definition

```yaml
{"operationId": "web3-hostname-list-ipfs-universal-path-gateway-content-list-entries", "summary": "List IPFS Universal Path Gateway Content List Entries", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "responses": {"200": {"description": "List IPFS Universal Path Gateway Content List Entries response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_content_list_entry_collection_response"}}}}, "4XX": {"description": "List IPFS Universal Path Gateway Content List Entries error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_entry_collection_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "List IPFS Universal Path Gateway Content List Entries response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_content_list_entry_collection_response"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-api-token-group": ["Web3 Hostnames Write", "Web3 Hostnames Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
