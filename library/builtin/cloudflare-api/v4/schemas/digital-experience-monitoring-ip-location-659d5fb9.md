---
title: digital-experience-monitoring_ip_location
page_id: schema-digital-experience-monitoring-ip-location-659d5fb9
path: schemas
description: Geographic location information. All fields are returned as the literal string `"REDACTED"` for callers that do not have the PII permission.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_ip_location

Geographic location information. All fields are returned as the literal string `"REDACTED"` for callers that do not have the PII permission.

```yaml
{"description": "Geographic location information. All fields are returned as the literal string `\"REDACTED\"` for callers that do not have the PII permission.\n", "type": "object", "properties": {"city": {"description": "City name. Returned as `\"REDACTED\"` without PII permission.", "type": "string", "example": "San Francisco"}, "country_iso": {"description": "Country ISO code. Returned as `\"REDACTED\"` without PII permission.", "type": "string", "example": "US"}, "state_iso": {"description": "State/province ISO code. Returned as `\"REDACTED\"` without PII permission.", "type": "string", "example": "CA"}, "zip": {"description": "ZIP/postal code. Returned as `\"REDACTED\"` without PII permission.", "type": "string", "example": "94107"}}}
```
