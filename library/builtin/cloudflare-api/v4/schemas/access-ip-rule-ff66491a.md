---
title: access_ip_rule
page_id: schema-access-ip-rule-ff66491a
path: schemas
description: Matches an IP address block.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_ip_rule

Matches an IP address block.

```yaml
{"description": "Matches an IP address block.", "type": "object", "properties": {"ip": {"type": "object", "properties": {"ip": {"description": "An IPv4 or IPv6 CIDR block.", "type": "string", "example": "2400:cb00:21:10a::/64"}}, "required": ["ip"]}}, "required": ["ip"], "title": "IP ranges"}
```
