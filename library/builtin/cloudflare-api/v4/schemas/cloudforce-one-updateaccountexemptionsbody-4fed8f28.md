---
title: cloudforce-one_UpdateAccountExemptionsBody
page_id: schema-cloudforce-one-updateaccountexemptionsbody-4fed8f28
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_UpdateAccountExemptionsBody

```yaml
{"type": "object", "properties": {"namespace": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_ExemptionUpdateEntry"}}, "tag_match": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_ExemptionUpdateEntry"}}, "worker_name": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_ExemptionUpdateEntry"}}}, "additionalProperties": false}
```
