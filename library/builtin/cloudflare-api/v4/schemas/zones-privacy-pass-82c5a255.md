---
title: zones_privacy_pass
page_id: schema-zones-privacy-pass-82c5a255
path: schemas
description: Privacy Pass v1 was a browser extension developed by the Privacy Pass Team to improve the browsing experience for your visitors by allowing users to reduce the number of CAPTCHAs shown. (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_privacy_pass

Privacy Pass v1 was a browser extension developed by the Privacy Pass Team to improve the browsing experience for your visitors by allowing users to reduce the number of CAPTCHAs shown. (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).

```yaml
{"description": "Privacy Pass v1 was a browser extension developed by the Privacy Pass Team to improve the browsing experience for your visitors by allowing users to reduce the number of CAPTCHAs shown. (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "privacy_pass", "enum": ["privacy_pass"]}, "value": {"$ref": "#/components/schemas/zones_privacy_pass_value"}}}], "deprecated": true, "title": "Privacy Pass v1 (deprecated)", "x-stainless-deprecation-message": "Privacy Pass v1 was deprecated in 2023. (Announcement - https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation details - https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)"}
```
