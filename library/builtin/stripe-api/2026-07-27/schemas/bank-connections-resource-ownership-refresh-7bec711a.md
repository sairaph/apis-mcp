---
title: bank_connections_resource_ownership_refresh
page_id: schema-bank-connections-resource-ownership-refresh-7bec711a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# bank_connections_resource_ownership_refresh

```yaml
{"title": "BankConnectionsResourceOwnershipRefresh", "required": ["last_attempted_at", "status"], "type": "object", "properties": {"last_attempted_at": {"type": "integer", "description": "The time at which the last refresh attempt was initiated. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "next_refresh_available_at": {"type": "integer", "description": "Time at which the next ownership refresh can be initiated. This value will be `null` when `status` is `pending`. Measured in seconds since the Unix epoch.", "format": "unix-time", "nullable": true}, "status": {"type": "string", "description": "The status of the last refresh attempt.", "enum": ["failed", "pending", "succeeded"]}}, "description": "", "x-expandableFields": []}
```
