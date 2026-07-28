---
title: zone-analytics-api_threats
page_id: schema-zone-analytics-api-threats-d0892691
path: schemas
description: Breakdown of totals for threats.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_threats

Breakdown of totals for threats.

```yaml
{"description": "Breakdown of totals for threats.", "type": "object", "properties": {"all": {"description": "The total number of identifiable threats received over the time frame.", "type": "integer"}, "country": {"description": "A list of key/value pairs where the key is a two-digit country code and the value is the number of malicious requests received from that country.", "type": "object", "example": {"AU": 91, "CN": 523423, "US": 123}}, "type": {"description": "The list of key/value pairs where the key is a threat category and the value is the number of requests.", "type": "object", "example": {"hot.ban.unknown": 5324, "macro.chl.captchaErr": 1341, "macro.chl.jschlErr": 5323, "user.ban.ip": 123}}}}
```
