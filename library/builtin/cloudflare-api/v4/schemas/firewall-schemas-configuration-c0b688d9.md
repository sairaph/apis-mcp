---
title: firewall_schemas-configuration
page_id: schema-firewall-schemas-configuration-c0b688d9
path: schemas
description: The configuration object for the current rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_schemas-configuration

The configuration object for the current rule.

```yaml
{"description": "The configuration object for the current rule.", "type": "object", "properties": {"target": {"description": "The configuration target for this rule. You must set the target to `ua` for User Agent Blocking rules.", "type": "string", "example": "ua"}, "value": {"description": "The exact user agent string to match. This value will be compared to the received `User-Agent` HTTP header value.", "type": "string", "example": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4"}}}
```
