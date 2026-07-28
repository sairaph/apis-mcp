---
title: digital-experience-monitoring_post_commands_response
page_id: schema-digital-experience-monitoring-post-commands-response-03970760
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_post_commands_response

```yaml
{"type": "object", "properties": {"commands": {"description": "List of created commands", "type": "array", "items": {"properties": {"args": {"description": "Command arguments", "type": "object", "additionalProperties": {"description": "Command argument value as a string", "type": "string"}}, "device_id": {"description": "Identifier for the device associated with the command", "type": "string"}, "id": {"description": "Unique identifier for the command", "type": "string"}, "registration_id": {"description": "Unique identifier for the device registration", "type": "string"}, "status": {"description": "Current status of the command", "type": "string", "enum": ["PENDING_EXEC", "PENDING_UPLOAD", "SUCCESS", "FAILED"]}, "type": {"description": "Type of the command (e.g., \"pcap\", \"speed-test\", or \"warp-diag\")", "type": "string"}}, "type": "object"}}}}
```
