---
title: zones_disable_security
page_id: schema-zones-disable-security-22eb59e6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_disable_security

```yaml
{"type": "object", "properties": {"id": {"description": "Turn off\n[Email Obfuscation](https://developers.cloudflare.com/waf/tools/scrape-shield/email-address-obfuscation/),\n[Rate Limiting (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-rate-limiting/),\n[Scrape Shield](https://developers.cloudflare.com/waf/tools/scrape-shield/),\n[URL (Zone) Lockdown](https://developers.cloudflare.com/waf/tools/zone-lockdown/), and\n[WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).\n", "type": "string", "enum": ["disable_security"], "x-auditable": true}}, "title": "Disable Security", "x-stainless-skip": ["terraform"]}
```
