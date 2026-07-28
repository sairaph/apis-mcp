---
title: Delete ACL
page_id: operation-delete-accounts-account-id-secondary-dns-acls-acl-id-474be2b2
path: operations/secondary-dns-acl
description: Delete ACL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/secondary_dns/acls/{acl_id}
operation_ids:
    - secondary-dns-(-acl)-delete-acl
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete ACL

`DELETE /accounts/{account_id}/secondary_dns/acls/{acl_id}`

Operation ID: `secondary-dns-(-acl)-delete-acl`

Delete ACL.

## Definition

```yaml
{"operationId": "secondary-dns-(-acl)-delete-acl", "summary": "Delete ACL", "description": "Delete ACL.", "parameters": [{"name": "acl_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete ACL response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_id_response-3"}}}}, "4XX": {"description": "Delete ACL response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_id_response-3"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (ACL)"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.acls", "x-fern-sdk-method-name": "delete"}
```
