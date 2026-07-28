---
title: zero-trust-gateway_dns_destination_ip_pair
page_id: schema-zero-trust-gateway-dns-destination-ip-pair-385a2b67
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_dns_destination_ip_pair

```yaml
{"type": "object", "properties": {"backup_ip": {"type": "string", "example": "172.64.36.2", "x-auditable": true}, "id": {"type": "string", "example": "0e4a32c6-6fb8-4858-9296-98f51631e8e6", "x-auditable": true}, "pair_type": {"description": "Specify whether the pair shared across multiple accounts (shared) or available exclusively to this account. Non-shared pairs can contain Cloudflare-owned IPs (dedicated) or customer-provided IPs (byoip).", "example": "shared", "enum": ["shared", "dedicated", "byoip"], "x-auditable": true}, "primary_ip": {"type": "string", "example": "172.64.36.1", "x-auditable": true}}, "required": ["id", "primary_ip", "backup_ip", "pair_type"]}
```
