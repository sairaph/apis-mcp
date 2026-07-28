---
title: email_sending_subdomain_preview_response
page_id: schema-email-sending-subdomain-preview-response-10562546
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_sending_subdomain_preview_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/email_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"errors": {"description": "DNS issues detected — missing records that will be created and conflicts with existing records.", "type": "array", "items": {"$ref": "#/components/schemas/email_sending_subdomain_config_error"}}, "records": {"description": "DNS records that would be created for the subdomain.", "type": "array", "items": {"$ref": "#/components/schemas/email_dns_record"}}}}}}]}
```
