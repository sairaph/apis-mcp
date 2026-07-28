---
title: ACL Details
page_id: operation-get-accounts-account-id-secondary-dns-acls-acl-id-cdd16c5a
path: operations/secondary-dns-acl
description: Get ACL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/secondary_dns/acls/{acl_id}
operation_ids:
    - secondary-dns-(-acl)-acl-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# ACL Details

`GET /accounts/{account_id}/secondary_dns/acls/{acl_id}`

Operation ID: `secondary-dns-(-acl)-acl-details`

Get ACL.

## Definition

```yaml
{"operationId": "secondary-dns-(-acl)-acl-details", "summary": "ACL Details", "description": "Get ACL.", "parameters": [{"name": "acl_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "responses": {"200": {"description": "ACL Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response-3"}}}}, "4XX": {"description": "ACL Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response-3"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (ACL)"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.acls", "x-fern-sdk-method-name": "get"}
```
