---
title: tls-certificates-and-hostnames_ct_alerting_emails
page_id: schema-tls-certificates-and-hostnames-ct-alerting-emails-31cf4e2b
path: schemas
description: Email addresses that receive CT alert notifications. Only present and configurable for Business and Enterprise zones. Maximum of 10 addresses. For Free and Pro zones, notifications are sent to all users with SSL permissions on the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_ct_alerting_emails

Email addresses that receive CT alert notifications. Only present and configurable for Business and Enterprise zones. Maximum of 10 addresses. For Free and Pro zones, notifications are sent to all users with SSL permissions on the zone.

```yaml
{"description": "Email addresses that receive CT alert notifications. Only present and configurable for Business and Enterprise zones. Maximum of 10 addresses. For Free and Pro zones, notifications are sent to all users with SSL permissions on the zone.\n", "type": "array", "items": {"format": "email", "type": "string"}, "example": ["security@example.com", "admin@example.com"], "maxItems": 10}
```
