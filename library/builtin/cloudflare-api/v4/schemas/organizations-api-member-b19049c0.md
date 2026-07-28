---
title: organizations-api_Member
page_id: schema-organizations-api-member-b19049c0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_Member

```yaml
{"type": "object", "properties": {"create_time": {"type": "string", "format": "date-time"}, "id": {"$ref": "#/components/schemas/organizations-api_MemberID"}, "meta": {"type": "object", "additionalProperties": {"type": "object"}}, "status": {"type": "string", "enum": ["active", "pending", "rejected", "canceled"]}, "update_time": {"type": "string", "format": "date-time"}, "user": {"$ref": "#/components/schemas/organizations-api_MemberSubjectUser"}}, "required": ["id", "status", "user", "meta", "create_time", "update_time"]}
```
