---
title: posture-api_DlpContext
page_id: schema-posture-api-dlpcontext-9dd1bdfc
path: schemas
description: DLP context information for a finding.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_DlpContext

DLP context information for a finding.

```yaml
{"description": "DLP context information for a finding.", "type": "object", "properties": {"created": {"description": "When the DLP context was created.", "type": "string", "format": "date-time", "example": "2025-03-18T17:25:38.695977Z", "readOnly": true}, "deleted": {"description": "When the DLP context was deleted.", "type": "string", "format": "date-time", "example": "2025-03-18T17:25:38.695977Z", "nullable": true}, "entry_ids": {"description": "DLP Entry IDs.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["21befc68-a297-4090-ac10-17a051b901cd", "d6dd1e16-f78c-401a-b564-45c4e44aa467"]}, "id": {"description": "Unique identifier for the DLP context.", "type": "string", "format": "uuid", "example": "7653ff3a-d25e-4c10-8034-3460937c045b"}, "match_context_max_extent": {"description": "DLP Right Boundary of match context.", "type": "integer", "example": 512, "maximum": 2147483647, "minimum": 0, "nullable": true}, "match_context_min_extent": {"description": "DLP Left Boundary of match context.", "type": "integer", "example": 1, "maximum": 2147483647, "minimum": 0, "nullable": true}, "match_context_payload": {"description": "DLP Match context payload that matched the profile in question.", "type": "object", "example": {}, "additionalProperties": true, "nullable": true}, "profile_id": {"description": "DLP Profile ID.", "type": "string", "format": "uuid", "example": "ab20a60b-21f2-4b13-ac98-24dcee27ac0e"}, "updated": {"description": "When the DLP context was last updated.", "type": "string", "format": "date-time", "example": "2025-03-18T17:25:38.695977Z", "readOnly": true}}, "required": ["created", "entry_ids", "profile_id", "updated"]}
```
