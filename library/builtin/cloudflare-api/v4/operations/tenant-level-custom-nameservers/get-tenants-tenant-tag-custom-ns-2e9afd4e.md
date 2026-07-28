---
title: List Tenant Custom Nameservers
page_id: operation-get-tenants-tenant-tag-custom-ns-e3922d8a
path: operations/tenant-level-custom-nameservers
description: List a tenant's custom nameservers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /tenants/{tenant_tag}/custom_ns
operation_ids:
    - tenant-level-custom-nameservers-list-tenant-custom-nameservers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Tenant Custom Nameservers

`GET /tenants/{tenant_tag}/custom_ns`

Operation ID: `tenant-level-custom-nameservers-list-tenant-custom-nameservers`

List a tenant's custom nameservers.

## Definition

```yaml
{"operationId": "tenant-level-custom-nameservers-list-tenant-custom-nameservers", "summary": "List Tenant Custom Nameservers", "description": "List a tenant's custom nameservers.", "parameters": [{"name": "tenant_tag", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier-3"}}], "responses": {"200": {"description": "List Tenant Custom Nameservers response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_tcns_response_collection"}}}}, "4XX": {"description": "List Tenant Custom Nameservers response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_tcns_response_collection"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Tenant-Level Custom Nameservers"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
