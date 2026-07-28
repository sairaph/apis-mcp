---
title: infra_Target
page_id: schema-infra-target-04ca292a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_Target

```yaml
{"type": "object", "properties": {"created_at": {"description": "Date and time at which the target was created", "type": "string", "format": "date-time", "example": "2019-08-24T14:15:22Z"}, "hostname": {"description": "A non-unique field that refers to a target", "type": "string", "example": "infra-access-target"}, "id": {"$ref": "#/components/schemas/infra_TargetId"}, "ip": {"$ref": "#/components/schemas/infra_IPInfo"}, "modified_at": {"description": "Date and time at which the target was modified", "type": "string", "format": "date-time", "example": "2019-08-24T14:15:22Z"}}, "required": ["id", "hostname", "ip", "created_at", "modified_at"]}
```
