---
title: posture-api_FindingInstanceBulkActionRequest
page_id: schema-posture-api-findinginstancebulkactionrequest-dab40e4f
path: schemas
description: Request body for bulk actions on finding instances.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingInstanceBulkActionRequest

Request body for bulk actions on finding instances.

```yaml
{"description": "Request body for bulk actions on finding instances.", "type": "object", "properties": {"check_instances": {"description": "A list of finding instance IDs to pass along.", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["497f6eca-6276-4993-bfeb-53cbbbba6f08"]}}, "required": ["check_instances"]}
```
