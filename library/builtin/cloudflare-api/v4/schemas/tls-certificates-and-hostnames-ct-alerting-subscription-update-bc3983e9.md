---
title: tls-certificates-and-hostnames_ct_alerting_subscription_update
page_id: schema-tls-certificates-and-hostnames-ct-alerting-subscription-update-bc3983e9
path: schemas
description: Request body for updating CT alerting subscription settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_ct_alerting_subscription_update

Request body for updating CT alerting subscription settings.

```yaml
{"description": "Request body for updating CT alerting subscription settings.", "type": "object", "properties": {"emails": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_emails"}, "enabled": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_ct_alerting_enabled"}}, "additionalProperties": false, "required": ["enabled"]}
```
