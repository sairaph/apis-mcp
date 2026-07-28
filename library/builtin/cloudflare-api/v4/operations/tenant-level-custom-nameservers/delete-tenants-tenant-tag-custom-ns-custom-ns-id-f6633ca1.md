---
title: Delete Tenant Custom Nameserver
page_id: operation-delete-tenants-tenant-tag-custom-ns-custom-ns-id-06a158b1
path: operations/tenant-level-custom-nameservers
description: Delete Tenant Custom Nameserver
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /tenants/{tenant_tag}/custom_ns/{custom_ns_id}
operation_ids:
    - tenant-level-custom-nameservers-delete-tenant-custom-nameserver
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Tenant Custom Nameserver

`DELETE /tenants/{tenant_tag}/custom_ns/{custom_ns_id}`

Operation ID: `tenant-level-custom-nameservers-delete-tenant-custom-nameserver`

## Definition

```yaml
{"operationId": "tenant-level-custom-nameservers-delete-tenant-custom-nameserver", "summary": "Delete Tenant Custom Nameserver", "parameters": [{"name": "custom_ns_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_ns_name"}}, {"name": "tenant_tag", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-custom-nameservers_identifier-3"}}], "responses": {"200": {"description": "Delete Tenant Custom Nameserver response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-custom-nameservers_empty_response"}}}}, "4XX": {"description": "Delete Tenant Custom Nameserver response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-custom-nameservers_empty_response"}, {"$ref": "#/components/schemas/dns-custom-nameservers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Tenant-Level Custom Nameservers"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
