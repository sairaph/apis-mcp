---
title: abuse-reports_MitigationListItem
page_id: schema-abuse-reports-mitigationlistitem-c3136320
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_MitigationListItem

```yaml
{"type": "object", "properties": {"effective_date": {"description": "Date when the mitigation will become active. Time in RFC 3339 format (https://www.rfc-editor.org/rfc/rfc3339.html)", "type": "string", "example": "2009-11-10T23:00:00Z"}, "entity_id": {"type": "string"}, "entity_type": {"$ref": "#/components/schemas/abuse-reports_MitigatedEntityType"}, "id": {"description": "ID of remediation.", "type": "string"}, "status": {"$ref": "#/components/schemas/abuse-reports_MitigationStatus"}, "type": {"$ref": "#/components/schemas/abuse-reports_MitigationType"}}, "required": ["id", "type", "effective_date", "status", "entity_type", "entity_id"]}
```
