---
title: Delete IPFS Universal Path Gateway Content List Entry
page_id: operation-delete-zones-zone-id-web3-hostnames-identifier-ipfs-universal-path-conte-0c8db018
path: operations/web3-hostname
description: Delete IPFS Universal Path Gateway Content List Entry
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries/{content_list_entry_identifier}
operation_ids:
    - web3-hostname-delete-ipfs-universal-path-gateway-content-list-entry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete IPFS Universal Path Gateway Content List Entry

`DELETE /zones/{zone_id}/web3/hostnames/{identifier}/ipfs_universal_path/content_list/entries/{content_list_entry_identifier}`

Operation ID: `web3-hostname-delete-ipfs-universal-path-gateway-content-list-entry`

## Definition

```yaml
{"operationId": "web3-hostname-delete-ipfs-universal-path-gateway-content-list-entry", "summary": "Delete IPFS Universal Path Gateway Content List Entry", "parameters": [{"name": "content_list_entry_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/web3_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete IPFS Universal Path Gateway Content List Entry response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/web3_api-response-single-id"}}}}, "4XX": {"description": "Delete IPFS Universal Path Gateway Content List Entry error response (4XX).", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_api-response-single-id"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}, "5XX": {"description": "Delete IPFS Universal Path Gateway Content List Entry response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/web3_api-response-single-id"}, {"$ref": "#/components/schemas/web3_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web3 Hostname"], "x-api-token-group": ["Web3 Hostnames Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
