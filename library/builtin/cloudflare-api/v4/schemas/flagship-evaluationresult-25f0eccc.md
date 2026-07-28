---
title: flagship_EvaluationResult
page_id: schema-flagship-evaluationresult-25f0eccc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# flagship_EvaluationResult

```yaml
{"type": "object", "properties": {"flagKey": {"type": "string"}, "reason": {"type": "string", "enum": ["TARGETING_MATCH", "DEFAULT", "DISABLED", "SPLIT"]}, "value": {"$ref": "#/components/schemas/flagship_JsonValue"}, "variant": {"type": "string"}}, "required": ["flagKey", "variant", "reason"]}
```
