---
title: zones_ssl_recommender
page_id: schema-zones-ssl-recommender-6e5fde28
path: schemas
description: Enrollment in the SSL/TLS Recommender service which tries to detect and recommend (by sending periodic emails) the most secure SSL/TLS setting your origin servers support.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_ssl_recommender

Enrollment in the SSL/TLS Recommender service which tries to detect and recommend (by sending periodic emails) the most secure SSL/TLS setting your origin servers support.

```yaml
{"description": "Enrollment in the SSL/TLS Recommender service which tries to detect and recommend (by sending periodic emails) the most secure SSL/TLS setting your origin servers support.", "allOf": [{"properties": {"enabled": {"$ref": "#/components/schemas/zones_ssl_recommender_enabled"}, "id": {"description": "Enrollment value for SSL/TLS Recommender.", "example": "ssl_recommender", "enum": ["ssl_recommender"]}}}], "title": "SSL/TLS Recommender"}
```
