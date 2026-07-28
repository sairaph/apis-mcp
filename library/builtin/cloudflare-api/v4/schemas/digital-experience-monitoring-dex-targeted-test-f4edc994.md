---
title: digital-experience-monitoring_dex-targeted-test
page_id: schema-digital-experience-monitoring-dex-targeted-test-f4edc994
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_dex-targeted-test

```yaml
{"type": "object", "properties": {"data": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_device-dex-test-schemas-data"}], "x-auditable": true}, "enabled": {"type": "boolean", "x-auditable": true}, "name": {"type": "string", "x-auditable": true}, "test_id": {"type": "string", "x-auditable": true}}, "required": ["test_id", "name", "enabled", "data"]}
```
