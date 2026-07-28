---
title: Add Tenant Custom Nameserver
page_id: operation-post-tenants-tenant-tag-custom-ns-23b4a4d2
path: operations/tenant-level-custom-nameservers
description: Add Tenant Custom Nameserver
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /tenants/{tenant_tag}/custom_ns
operation_ids:
    - tenant-level-custom-nameservers-add-tenant-custom-nameserver
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add Tenant Custom Nameserver

`POST /tenants/{tenant_tag}/custom_ns`

Operation ID: `tenant-level-custom-nameservers-add-tenant-custom-nameserver`

## Definition

```yaml
{"operationId": "tenant-level-custom-nameservers-add-tenant-custom-nameserver", "summary": "Add Tenant Custom Nameserver", "parameters": [{"name": "tenant_tag", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier-3"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_CustomNSInput"}}}}, "responses": {"200": {"description": "Add Tenant Custom Nameserver response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_tcns_response_single"}}}}, "4XX": {"description": "Add Tenant Custom Nameserver response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_tcns_response_single"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Tenant-Level Custom Nameservers"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
