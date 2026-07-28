---
title: zones_tls_client_auth
page_id: schema-zones-tls-client-auth-d6a5eb50
path: schemas
description: TLS Client Auth requires Cloudflare to connect to your origin server using a client certificate (Enterprise Only).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_tls_client_auth

TLS Client Auth requires Cloudflare to connect to your origin server using a client certificate (Enterprise Only).

```yaml
{"description": "TLS Client Auth requires Cloudflare to connect to your origin server using a client certificate (Enterprise Only).", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "tls_client_auth", "enum": ["tls_client_auth"]}, "value": {"$ref": "#/components/schemas/zones_tls_client_auth_value"}}}], "title": "TLS Client Authentication"}
```
