---
title: security-center_scanStatusResponse
page_id: schema-security-center-scanstatusresponse-0295c3ab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_scanStatusResponse

```yaml
{"type": "object", "properties": {"scan_id": {"description": "An opaque identifier for the scan.", "type": "string", "example": "d5e94e48-504f-4a7f-a8c4-e0dc2e05e5f2"}, "started_at": {"description": "The time at which the scan was started, in RFC 3339 format.", "type": "string", "format": "date-time", "example": "2026-04-13T23:59:59Z"}, "status": {"description": "The current status of the scan.", "type": "string", "enum": ["in_progress", "completed"]}}, "required": ["scan_id", "started_at", "status"]}
```
