---
title: zones_security_header_value
page_id: schema-zones-security-header-value-9c503853
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_security_header_value

```yaml
{"type": "object", "properties": {"strict_transport_security": {"description": "Strict Transport Security.", "type": "object", "properties": {"enabled": {"description": "Whether or not strict transport security is enabled.", "type": "boolean", "example": true}, "include_subdomains": {"description": "Include all subdomains for strict transport security.", "type": "boolean", "example": true}, "max_age": {"description": "Max age in seconds of the strict transport security.", "type": "number", "example": 86400}, "nosniff": {"description": "Whether or not to include 'X-Content-Type-Options: nosniff' header.", "type": "boolean", "example": true}, "preload": {"description": "Enable automatic preload of the HSTS configuration.", "type": "boolean", "example": true}}}}, "default": {"strict_transport_security": {"enabled": true, "include_subdomains": true, "max_age": 86400, "nosniff": true, "preload": false}}}
```
