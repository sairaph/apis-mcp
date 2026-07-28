---
title: digital-experience-monitoring_get_commands_quota_response
page_id: schema-digital-experience-monitoring-get-commands-quota-response-7031bf1f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_get_commands_quota_response

```yaml
{"type": "object", "properties": {"quota": {"description": "The total number of commands that can be initiated for an account.", "type": "number"}, "quota_usage": {"description": "The number of commands that have been initiated for an account.", "type": "number"}, "reset_time": {"description": "The time when the quota resets.", "type": "string", "format": "date-time"}}, "required": ["quota_usage", "quota", "reset_time"]}
```
