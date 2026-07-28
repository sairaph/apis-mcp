---
title: security-center_userClassificationUpdate
page_id: schema-security-center-userclassificationupdate-f5cad4e6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# security-center_userClassificationUpdate

```yaml
{"type": "object", "properties": {"classification": {"$ref": "#/components/schemas/security-center_userClassification"}, "rationale": {"description": "Rationale for the classification change. Required when classification is 'accept_risk' or 'other'.", "type": "string", "x-auditable": true}}}
```
