---
title: List ACLs
page_id: operation-get-accounts-account-id-secondary-dns-acls-fa8ca8ea
path: operations/secondary-dns-acl
description: List ACLs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secondary_dns/acls
operation_ids:
    - secondary-dns-(-acl)-list-ac-ls
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List ACLs

`GET /accounts/{account_id}/secondary_dns/acls`

Operation ID: `secondary-dns-(-acl)-list-ac-ls`

List ACLs.

## Definition

```yaml
{"operationId": "secondary-dns-(-acl)-list-ac-ls", "summary": "List ACLs", "description": "List ACLs.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "responses": {"200": {"description": "List ACLs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_response_collection-3"}}}}, "4XX": {"description": "List ACLs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_response_collection-3"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (ACL)"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.acls", "x-fern-sdk-method-name": "list"}
```
