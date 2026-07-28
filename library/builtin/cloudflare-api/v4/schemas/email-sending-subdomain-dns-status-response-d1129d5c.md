---
title: email_sending_subdomain_dns_status_response
page_id: schema-email-sending-subdomain-dns-status-response-d1129d5c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_sending_subdomain_dns_status_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/email_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"errors": {"description": "DNS issues detected against the current zone state.", "type": "array", "items": {"$ref": "#/components/schemas/email_sending_subdomain_config_error"}}, "records": {"description": "Desired DNS records for the subdomain.", "type": "array", "items": {"$ref": "#/components/schemas/email_dns_record"}}, "status": {"description": "Aggregated DNS state for the subdomain. `unlocked` means desired records exist with correct content but at least one has had its email_routing/read_only lock cleared.", "type": "string", "enum": ["ready", "unconfigured", "unlocked", "misconfigured"]}}}}}]}
```
