---
title: workers_ErrorWorkerNameSubdomainLengthLimit
page_id: schema-workers-errorworkernamesubdomainlengthlimit-c9670366
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerNameSubdomainLengthLimit

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the Worker name is too long to be used as a subdomain.", "type": "integer", "enum": [100132]}, "message": {"description": "Message explaining that the Worker name exceeds the 63 character limit for subdomains.", "type": "string"}}, "required": ["code", "message"]}
```
