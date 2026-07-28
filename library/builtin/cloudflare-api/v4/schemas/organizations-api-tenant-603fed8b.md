---
title: organizations-api_Tenant
page_id: schema-organizations-api-tenant-603fed8b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_Tenant

```yaml
{"type": "object", "properties": {"cdate": {"type": "string", "format": "date-time"}, "customer_id": {"type": "string"}, "edate": {"type": "string", "format": "date-time"}, "tenant_contacts": {"type": "object", "properties": {"email": {"type": "string"}, "website": {"type": "string"}}}, "tenant_labels": {"type": "array", "items": {"type": "string"}}, "tenant_metadata": {"type": "object", "properties": {"dns": {"type": "object", "properties": {"ns_pool": {"type": "object", "properties": {"primary": {"type": "string"}, "secondary": {"type": "string"}}}}, "required": ["ns_pool"]}}}, "tenant_name": {"type": "string"}, "tenant_network": {"type": "object"}, "tenant_status": {"type": "string"}, "tenant_tag": {"type": "string"}, "tenant_type": {"type": "string"}, "tenant_units": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_TenantUnit"}}}, "required": ["tenant_tag", "tenant_name", "tenant_labels", "tenant_type", "tenant_status", "tenant_metadata", "tenant_contacts", "cdate", "edate", "tenant_network", "tenant_units"]}
```
