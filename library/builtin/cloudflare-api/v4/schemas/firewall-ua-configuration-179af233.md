---
title: firewall_ua_configuration
page_id: schema-firewall-ua-configuration-179af233
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_ua_configuration

```yaml
{"type": "object", "properties": {"target": {"description": "The configuration target. You must set the target to `ua` when specifying a user agent in the rule.", "type": "string", "example": "ua", "enum": ["ua"]}, "value": {"description": "the user agent to exactly match", "type": "string", "example": "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)", "x-auditable": true}}, "title": "A user agent configuration."}
```
