---
title: aaa_audit-logs-v2-org
page_id: schema-aaa-audit-logs-v2-org-6dfaf43c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-logs-v2-org

```yaml
{"type": "object", "properties": {"action": {"$ref": "#/components/schemas/aaa_audit-log-action"}, "actor": {"allOf": [{"$ref": "#/components/schemas/aaa_audit-log-actor-base"}, {"properties": {"type": {"description": "The type of actor.", "type": "string", "example": "user", "enum": ["cloudflare_admin", "system", "user"]}}, "type": "object"}]}, "id": {"$ref": "#/components/schemas/aaa_identifier-2"}, "organization": {"description": "Contains organization related information.", "type": "object", "properties": {"id": {"description": "A unique identifier for the organization.", "type": "string", "example": "019c4f65e7607d8c9f6f6b58aa3aff50"}}}, "raw": {"$ref": "#/components/schemas/aaa_audit-log-raw"}, "resource": {"description": "Provides details about the affected resource.", "type": "object", "properties": {"id": {"description": "The unique identifier for the affected resource.", "type": "string"}, "product": {"description": "The Cloudflare product associated with the resource.", "type": "string", "example": "organizations"}, "request": {"type": "object"}, "response": {"type": "object"}, "scope": {"description": "The scope of the resource.", "type": "object", "example": "organizations"}, "type": {"description": "The type of the resource.", "type": "string"}}}}}
```
