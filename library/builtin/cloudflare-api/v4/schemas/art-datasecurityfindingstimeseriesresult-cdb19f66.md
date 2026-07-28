---
title: art_DataSecurityFindingsTimeseriesResult
page_id: schema-art-datasecurityfindingstimeseriesresult-cdb19f66
path: schemas
description: Merged CASB and CDE findings timeseries result.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_DataSecurityFindingsTimeseriesResult

Merged CASB and CDE findings timeseries result.

```yaml
{"description": "Merged CASB and CDE findings timeseries result.", "type": "object", "properties": {"resolution": {"description": "Always null for this endpoint.", "type": "string", "nullable": true}, "slots": {"description": "Contains time-bucketed result rows. Each slot includes a `timestamp` plus `content` and `posture` maps with `cloud` and `saas` keys.\n", "type": "array", "items": {"$ref": "#/components/schemas/art_ResultValues"}, "example": [{"content": {"cloud": 150, "saas": 23}, "posture": {"cloud": 0, "saas": 5}, "timestamp": "2024-11-05T00:00:00Z"}]}}, "required": ["slots"]}
```
