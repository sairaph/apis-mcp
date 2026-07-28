---
title: zone-analytics-api_since
page_id: schema-zone-analytics-api-since-088aa41c
path: schemas
description: |-
    The (inclusive) beginning of the requested time frame. This value can be a negative integer representing the number of minutes in the past relative to time the request is made, or can be an absolute timestamp that conforms to RFC 3339. At this point in time, it cannot exceed a time in the past greater than one year.

    Ranges that the Cloudflare web application provides will provide the following period length for each point:
    - Last 60 minutes (from -59 to -1): 1 minute resolution
    - Last 7 hours (from -419 to -60): 15 minutes resolution
    - Last 15 hours (from -899 to -420): 30 minutes resolution
    - Last 72 hours (from -4320 to -900): 1 hour resolution
    - Older than 3 days (-525600 to -4320): 1 day resolution.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_since

The (inclusive) beginning of the requested time frame. This value can be a negative integer representing the number of minutes in the past relative to time the request is made, or can be an absolute timestamp that conforms to RFC 3339. At this point in time, it cannot exceed a time in the past greater than one year.

Ranges that the Cloudflare web application provides will provide the following period length for each point:
- Last 60 minutes (from -59 to -1): 1 minute resolution
- Last 7 hours (from -419 to -60): 15 minutes resolution
- Last 15 hours (from -899 to -420): 30 minutes resolution
- Last 72 hours (from -4320 to -900): 1 hour resolution
- Older than 3 days (-525600 to -4320): 1 day resolution.

```yaml
{"description": "The (inclusive) beginning of the requested time frame. This value can be a negative integer representing the number of minutes in the past relative to time the request is made, or can be an absolute timestamp that conforms to RFC 3339. At this point in time, it cannot exceed a time in the past greater than one year.\n\nRanges that the Cloudflare web application provides will provide the following period length for each point:\n- Last 60 minutes (from -59 to -1): 1 minute resolution\n- Last 7 hours (from -419 to -60): 15 minutes resolution\n- Last 15 hours (from -899 to -420): 30 minutes resolution\n- Last 72 hours (from -4320 to -900): 1 hour resolution\n- Older than 3 days (-525600 to -4320): 1 day resolution.", "example": "2015-01-01T12:23:00Z", "default": -10080, "anyOf": [{"type": "string"}, {"type": "integer"}]}
```
