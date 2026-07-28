---
title: aaa_audit-logs-v2-history-result-info
page_id: schema-aaa-audit-logs-v2-history-result-info-e9212a50
path: schemas
description: Provides information about the result of the request, including count, cursor, and identification quality.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-logs-v2-history-result-info

Provides information about the result of the request, including count, cursor, and identification quality.

```yaml
{"description": "Provides information about the result of the request, including count, cursor, and identification quality.", "type": "object", "properties": {"count": {"description": "The number of records returned in the response.", "type": "integer", "example": 1}, "cursor": {"description": "The cursor token used for pagination.", "type": "string", "example": "ASqdKd7dKgxh-aZ8bm0mZos1BtW4BdEqifCzNkEeGRzi_5SN_-362Y8sF-C1TRn60_6rd3z2dIajf9EAPyQ_NmIeAMkacmaJPXipqvP7PLU4t72wyqBeJfjmjdE="}, "history_status": {"$ref": "#/components/schemas/aaa_history-status"}}, "required": ["count", "history_status"]}
```
