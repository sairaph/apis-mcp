---
title: digital-experience-monitoring_get_commands_response
page_id: schema-digital-experience-monitoring-get-commands-response-dd88e745
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_get_commands_response

```yaml
{"type": "object", "properties": {"commands": {"type": "array", "items": {"properties": {"completed_date": {"type": "string", "format": "date-time", "nullable": true}, "created_date": {"type": "string", "format": "date-time"}, "device_id": {"type": "string"}, "filename": {"type": "string", "nullable": true}, "id": {"type": "string"}, "registration_id": {"description": "Unique identifier for the device registration", "type": "string"}, "status": {"type": "string"}, "type": {"type": "string"}, "user_email": {"type": "string"}}, "type": "object"}}}}
```
