---
title: rlanalytics_rate_limit_analytics
page_id: schema-rlanalytics-rate-limit-analytics-ed7cbeaf
path: schemas
description: Rate limiting analytics for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rlanalytics_rate_limit_analytics

Rate limiting analytics for a zone.

```yaml
{"description": "Rate limiting analytics for a zone.", "type": "object", "properties": {"labels": {"description": "Mapping from rule ID to human-readable description of the rule.", "type": "object", "example": {"123e4567-e89b-12d3-a456-426655440000": "block login brute force on /login", "66698345-bc87-d312-a456-937482934748": "rate limit whole site"}, "additionalProperties": {"type": "string"}}, "since": {"description": "Start of the queried time period formatted as RFC 3339.", "type": "string", "format": "date-time", "example": "2024-01-01T00:00:00Z"}, "time_delta": {"description": "Length (in seconds) of the time segments dividing the entire time period.", "type": "integer", "example": 3600, "enum": [60, 3600, 86400, 2592000]}, "timeseries": {"description": "Time series with analytics data for each time segment.", "type": "array", "items": {"$ref": "#/components/schemas/rlanalytics_rate_limit_analytics_timeseries_point"}}, "until": {"description": "Exclusive end of the queried time period formatted as RFC 3339.", "type": "string", "format": "date-time", "example": "2024-01-02T00:00:00Z"}, "zone_id": {"description": "Numeric ID of the zone.", "type": "integer", "format": "int64", "example": 123}}, "required": ["since", "until", "time_delta", "labels", "timeseries"]}
```
