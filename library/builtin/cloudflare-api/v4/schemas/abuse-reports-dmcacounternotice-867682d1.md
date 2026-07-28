---
title: abuse-reports_DMCACounterNotice
page_id: schema-abuse-reports-dmcacounternotice-867682d1
path: schemas
description: Counter-notice details supporting an appeal.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_DMCACounterNotice

Counter-notice details supporting an appeal.

```yaml
{"description": "Counter-notice details supporting an appeal.", "type": "object", "properties": {"city": {"type": "string"}, "company": {"type": "string"}, "counter_notice_response": {"type": "string"}, "country": {"type": "string"}, "email": {"type": "string", "format": "email"}, "full_name": {"type": "string"}, "jurisdiction_consent": {"type": "boolean"}, "perjury_attestation": {"type": "boolean"}, "phone_number": {"type": "string"}, "signature": {"type": "string"}, "state": {"type": "string"}, "street_address": {"type": "string"}, "urls": {"type": "array", "items": {"format": "uri", "type": "string"}}, "zip_code": {"type": "string"}}, "required": ["full_name", "street_address", "city", "state", "country", "zip_code", "phone_number", "email", "urls", "perjury_attestation", "jurisdiction_consent", "signature"]}
```
