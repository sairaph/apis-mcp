---
title: zones_true_client_ip_header
page_id: schema-zones-true-client-ip-header-17624620
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_true_client_ip_header

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off the True-Client-IP Header feature of the Cloudflare Network app.\n", "type": "string", "example": "true_client_ip_header", "enum": ["true_client_ip_header"], "x-auditable": true}, "value": {"description": "The status of True Client IP Header.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "True Client IP Header", "x-stainless-skip": ["terraform"]}
```
