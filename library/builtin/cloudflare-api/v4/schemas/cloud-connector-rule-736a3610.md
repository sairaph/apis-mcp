---
title: cloud-connector_rule
page_id: schema-cloud-connector-rule-736a3610
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloud-connector_rule

```yaml
{"type": "object", "properties": {"description": {"type": "string", "example": "Rule description", "x-auditable": true}, "enabled": {"type": "boolean", "example": true, "x-auditable": true}, "expression": {"type": "string", "example": "http.cookie eq \"a=b\"", "x-auditable": true}, "id": {"type": "string", "example": "95c365e17e1b46599cd99e5b231fac4e", "x-auditable": true}, "parameters": {"description": "Parameters of Cloud Connector Rule", "properties": {"host": {"description": "Host to perform Cloud Connection to", "type": "string", "example": "examplebucket.s3.eu-north-1.amazonaws.com", "x-auditable": true}}, "type": "object"}, "provider": {"$ref": "#/components/schemas/cloud-connector_provider"}}}
```
