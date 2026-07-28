---
title: zones_redirects_for_ai_training
page_id: schema-zones-redirects-for-ai-training-cb008847
path: schemas
description: |-
    When enabled, Cloudflare will redirect verified AI training crawlers to canonical URLs
    found in the HTML response, ensuring AI models train on authoritative content.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_redirects_for_ai_training

When enabled, Cloudflare will redirect verified AI training crawlers to canonical URLs
found in the HTML response, ensuring AI models train on authoritative content.

```yaml
{"description": "When enabled, Cloudflare will redirect verified AI training crawlers to canonical URLs\nfound in the HTML response, ensuring AI models train on authoritative content.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "redirects_for_ai_training", "enum": ["redirects_for_ai_training"]}, "value": {"$ref": "#/components/schemas/zones_redirects_for_ai_training_value"}}}], "title": "Redirects for AI Training"}
```
