---
title: zero-trust-gateway_generate-cert-request
page_id: schema-zero-trust-gateway-generate-cert-request-37473d39
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_generate-cert-request

```yaml
{"type": "object", "properties": {"validity_period_days": {"description": "Sets the certificate validity period in days (range: 1-10,950 days / ~30 years). Defaults to 1,825 days (5 years). **Important**: This field is only settable during the certificate creation.  Certificates becomes immutable after creation - use the `/activate` and `/deactivate` endpoints to manage certificate lifecycle.", "type": "integer", "example": 1826, "x-auditable": true}}}
```
