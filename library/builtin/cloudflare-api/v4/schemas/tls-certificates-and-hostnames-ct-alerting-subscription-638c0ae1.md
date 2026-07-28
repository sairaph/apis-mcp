---
title: tls-certificates-and-hostnames_ct_alerting_subscription
page_id: schema-tls-certificates-and-hostnames-ct-alerting-subscription-638c0ae1
path: schemas
description: Certificate Transparency alerting subscription settings for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_ct_alerting_subscription

Certificate Transparency alerting subscription settings for a zone.

```yaml
{"description": "Certificate Transparency alerting subscription settings for a zone.", "type": "object", "properties": {"emails": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_emails"}, "enabled": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_enabled"}}, "required": ["enabled"]}
```
