---
title: email-security_PatternType
page_id: schema-email-security-patterntype-847cd986
path: schemas
description: |-
    Type of pattern matching.
    - EMAIL: matches a full email address (e.g. `user@example.com`)
    - DOMAIN: matches a domain name (e.g. `example.com`)
    - IP: matches a plain IPv4 address (e.g. `1.2.3.4`) or an IPv4 CIDR block (e.g. `1.2.3.0/24`). The API accepts only globally reachable addresses.
    - UNKNOWN: deprecated; you cannot use this when creating or updating policies, but it may appear on existing entries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_PatternType

Type of pattern matching.
- EMAIL: matches a full email address (e.g. `user@example.com`)
- DOMAIN: matches a domain name (e.g. `example.com`)
- IP: matches a plain IPv4 address (e.g. `1.2.3.4`) or an IPv4 CIDR block (e.g. `1.2.3.0/24`). The API accepts only globally reachable addresses.
- UNKNOWN: deprecated; you cannot use this when creating or updating policies, but it may appear on existing entries.

```yaml
{"description": "Type of pattern matching.\n- EMAIL: matches a full email address (e.g. `user@example.com`)\n- DOMAIN: matches a domain name (e.g. `example.com`)\n- IP: matches a plain IPv4 address (e.g. `1.2.3.4`) or an IPv4 CIDR block (e.g. `1.2.3.0/24`). The API accepts only globally reachable addresses.\n- UNKNOWN: deprecated; you cannot use this when creating or updating policies, but it may appear on existing entries.\n", "type": "string", "example": "EMAIL", "enum": ["EMAIL", "DOMAIN", "IP", "UNKNOWN"], "x-auditable": true}
```
