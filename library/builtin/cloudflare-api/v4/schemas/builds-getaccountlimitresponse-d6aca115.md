---
title: builds_GetAccountLimitResponse
page_id: schema-builds-getaccountlimitresponse-d6aca115
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_GetAccountLimitResponse

```yaml
{"type": "object", "properties": {"build_minutes_refresh_on": {"description": "When build minutes will refresh (only for non-paid plans)", "type": "string", "format": "date-time", "nullable": true}, "has_reached_build_minutes_limit": {"description": "Whether build minutes limit has been reached (only for non-paid plans)", "type": "boolean", "nullable": true}}}
```
