---
title: Create ACL
page_id: operation-post-accounts-account-id-secondary-dns-acls-6675997e
path: operations/secondary-dns-acl
description: Create ACL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/secondary_dns/acls
operation_ids:
    - secondary-dns-(-acl)-create-acl
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create ACL

`POST /accounts/{account_id}/secondary_dns/acls`

Operation ID: `secondary-dns-(-acl)-create-acl`

Create ACL.

## Definition

```yaml
{"operationId": "secondary-dns-(-acl)-create-acl", "summary": "Create ACL", "description": "Create ACL.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"ip_range": {"$ref": "#/components/schemas/secondary-dns_ip_range"}, "name": {"$ref": "#/components/schemas/secondary-dns_name-4"}}, "required": ["name", "ip_range"]}}}}, "responses": {"200": {"description": "Create ACL response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response-3"}}}}, "4XX": {"description": "Create ACL response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response-3"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (ACL)"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.acls", "x-fern-sdk-method-name": "create"}
```
