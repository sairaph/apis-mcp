---
title: email_account_rules_response_collection
page_id: schema-email-account-rules-response-collection-670fb82b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_account_rules_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/email_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/email_account_rule"}}, "result_info": {"type": "object", "properties": {"count": {"example": 1}, "page": {"example": 1}, "per_page": {"example": 20}, "total_count": {"example": 1}}}}}]}
```
