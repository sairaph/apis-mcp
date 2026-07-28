---
title: logshare_start
page_id: schema-logshare-start-5d105664
path: schemas
description: Sets the (inclusive) beginning of the requested time frame. This can be a unix timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to RFC 3339. At this point in time, it cannot exceed a time in the past greater than seven days.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logshare_start

Sets the (inclusive) beginning of the requested time frame. This can be a unix timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to RFC 3339. At this point in time, it cannot exceed a time in the past greater than seven days.

```yaml
{"description": "Sets the (inclusive) beginning of the requested time frame. This can be a unix timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to RFC 3339. At this point in time, it cannot exceed a time in the past greater than seven days.", "example": "2018-05-20T10:00:00Z", "anyOf": [{"type": "string"}, {"type": "integer"}], "x-auditable": true}
```
