---
title: zone-analytics-api_query_response
page_id: schema-zone-analytics-api-query-response-2a9da5c4
path: schemas
description: The exact parameters/timestamps the analytics service used to return data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_query_response

The exact parameters/timestamps the analytics service used to return data.

```yaml
{"description": "The exact parameters/timestamps the analytics service used to return data.", "type": "object", "properties": {"since": {"$ref": "#/components/schemas/zone-analytics-api_since"}, "time_delta": {"description": "The amount of time (in minutes) that each data point in the timeseries represents. The granularity of the time-series returned (e.g. each bucket in the time series representing 1-minute vs 1-day) is calculated by the API based on the time-range provided to the API.", "type": "integer"}, "until": {"$ref": "#/components/schemas/zone-analytics-api_until"}}, "readOnly": true}
```
