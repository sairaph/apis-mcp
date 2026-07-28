---
title: stream_accessRules
page_id: schema-stream-accessrules-293ce7af
path: schemas
description: Defines rules for fine-grained control over content than signed URL tokens alone. Access rules primarily make tokens conditionally valid based on user information. Access Rules are specified on token payloads as the `accessRules` property containing an array of Rule objects.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_accessRules

Defines rules for fine-grained control over content than signed URL tokens alone. Access rules primarily make tokens conditionally valid based on user information. Access Rules are specified on token payloads as the `accessRules` property containing an array of Rule objects.

```yaml
{"description": "Defines rules for fine-grained control over content than signed URL tokens alone. Access rules primarily make tokens conditionally valid based on user information. Access Rules are specified on token payloads as the `accessRules` property containing an array of Rule objects.", "type": "object", "properties": {"action": {"description": "The action to take when a request matches a rule. If the action is `block`, the signed token blocks views for viewers matching the rule.", "type": "string", "example": "allow", "enum": ["allow", "block"], "x-auditable": true}, "country": {"description": "An array of 2-letter country codes in ISO 3166-1 Alpha-2 format used to match requests.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "ip": {"description": "An array of IPv4 or IPV6 addresses or CIDRs used to match requests.", "type": "array", "items": {"type": "string", "x-auditable": true}}, "type": {"description": "Lists available rule types to match for requests. An `any` type matches all requests and can be used as a wildcard to apply default actions after other rules.", "type": "string", "example": "ip.src", "enum": ["any", "ip.src", "ip.geoip.country"], "x-auditable": true}}}
```
