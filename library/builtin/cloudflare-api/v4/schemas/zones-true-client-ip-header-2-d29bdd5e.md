---
title: zones_true_client_ip_header-2
page_id: schema-zones-true-client-ip-header-2-d29bdd5e
path: schemas
description: Allows customer to continue to use True Client IP (Akamai feature) in the headers we send to the origin. This is limited to Enterprise Zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_true_client_ip_header-2

Allows customer to continue to use True Client IP (Akamai feature) in the headers we send to the origin. This is limited to Enterprise Zones.

```yaml
{"description": "Allows customer to continue to use True Client IP (Akamai feature) in the headers we send to the origin. This is limited to Enterprise Zones.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "true_client_ip_header", "enum": ["true_client_ip_header"]}, "value": {"$ref": "#/components/schemas/zones_true_client_ip_header_value"}}}], "title": "True Client IP Header"}
```
