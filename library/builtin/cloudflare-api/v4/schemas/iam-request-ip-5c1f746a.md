---
title: iam_request_ip
page_id: schema-iam-request-ip-5c1f746a
path: schemas
description: Client IP restrictions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_request_ip

Client IP restrictions.

```yaml
{"description": "Client IP restrictions.", "type": "object", "properties": {"in": {"$ref": "#/components/schemas/iam_cidr_list"}, "not_in": {"$ref": "#/components/schemas/iam_cidr_list"}}, "example": {"in": ["123.123.123.0/24", "2606:4700::/32"], "not_in": ["123.123.123.100/24", "2606:4700:4700::/48"]}}
```
