---
title: email_account_rule
page_id: schema-email-account-rule-2d712c5d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_account_rule

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/email_rule_properties"}, {"properties": {"zone": {"description": "Zone information for the routing rule.", "type": "object", "properties": {"name": {"description": "Zone name.", "type": "string", "example": "example.com", "readOnly": true}, "tag": {"description": "Zone tag.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "readOnly": true}}}}}]}
```
