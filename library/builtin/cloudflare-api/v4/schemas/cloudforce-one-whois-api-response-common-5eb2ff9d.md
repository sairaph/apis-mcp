---
title: cloudforce-one-whois_api-response-common
page_id: schema-cloudforce-one-whois-api-response-common-5eb2ff9d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-whois_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/cloudforce-one-whois_schemas-messages"}, "messages": {"$ref": "#/components/schemas/cloudforce-one-whois_schemas-messages"}, "success": {"description": "Returns a boolean for the success/failure of the API call.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
