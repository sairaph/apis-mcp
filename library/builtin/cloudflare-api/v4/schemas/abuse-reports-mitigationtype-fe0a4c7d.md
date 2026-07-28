---
title: abuse-reports_MitigationType
page_id: schema-abuse-reports-mitigationtype-fe0a4c7d
path: schemas
description: The type of mitigation applied to a reported entity.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_MitigationType

The type of mitigation applied to a reported entity.

```yaml
{"description": "The type of mitigation applied to a reported entity.", "type": "string", "enum": ["account_suspend", "copyright_interstitial", "geo_block", "legal_block", "malware_interstitial", "misleading_interstitial", "network_block", "phishing_interstitial", "playfairite_enforce", "r2_takedown_account", "r2_takedown_bucket", "r2_takedown_object", "rate_limit_cache", "redirect_video_stream", "zone_fint", "registrar_freeze", "registrar_parking", "stream_block_account", "user_suspend", "workers_takedown_by_zone_id"]}
```
