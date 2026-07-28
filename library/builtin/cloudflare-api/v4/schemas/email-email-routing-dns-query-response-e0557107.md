---
title: email_email_routing_dns_query_response
page_id: schema-email-email-routing-dns-query-response-e0557107
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_email_routing_dns_query_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/email_api-response-collection"}, {"properties": {"result": {"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/email_email_routing_get_response_dns_errors"}, "record": {"type": "array", "items": {"$ref": "#/components/schemas/email_dns_record"}}}}}}]}
```
