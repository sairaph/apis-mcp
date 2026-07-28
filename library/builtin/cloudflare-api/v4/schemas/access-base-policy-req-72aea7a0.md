---
title: access_base_policy_req
page_id: schema-access-base-policy-req-72aea7a0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_base_policy_req

```yaml
{"type": "object", "properties": {"decision": {"$ref": "#/components/schemas/access_decision"}, "exclude": {"$ref": "#/components/schemas/access_exclude-2"}, "include": {"$ref": "#/components/schemas/access_include-2"}, "name": {"$ref": "#/components/schemas/access_name-9"}, "require": {"$ref": "#/components/schemas/access_require-2"}}, "required": ["name", "decision", "include"]}
```
