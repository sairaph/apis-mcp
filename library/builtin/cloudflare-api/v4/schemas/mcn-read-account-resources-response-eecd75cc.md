---
title: mcn_read_account_resources_response
page_id: schema-mcn-read-account-resources-response-eecd75cc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_read_account_resources_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/mcn_good_response_collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_details"}}, "result_info": {"$ref": "#/components/schemas/mcn_result_info"}}, "type": "object"}]}
```
