---
title: zero-trust-gateway_expiration
page_id: schema-zero-trust-gateway-expiration-3686d34a
path: schemas
description: Defines the expiration time stamp and default duration of a DNS policy. Takes precedence over the policy's `schedule` configuration, if any. This  does not apply to HTTP or network policies. Settable only for `dns` rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_expiration

Defines the expiration time stamp and default duration of a DNS policy. Takes precedence over the policy's `schedule` configuration, if any. This  does not apply to HTTP or network policies. Settable only for `dns` rules.

```yaml
{"description": "Defines the expiration time stamp and default duration of a DNS policy. Takes precedence over the policy's `schedule` configuration, if any. This  does not apply to HTTP or network policies. Settable only for `dns` rules.", "type": "object", "properties": {"duration": {"description": "Defines the default duration a policy active in minutes. Must set in order to use the `reset_expiration` endpoint on this rule.", "type": "integer", "example": 10, "minimum": 5, "x-auditable": true}, "expired": {"description": "Indicates whether the policy is expired.", "type": "boolean", "example": false, "readOnly": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed"}, "expires_at": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_timestamp"}, {"description": "Show the timestamp when the policy expires and stops applying.  The value must follow RFC 3339 and include a UTC offset.  The system accepts non-zero offsets but converts them to the equivalent UTC+00:00  value and returns timestamps with a trailing Z. Expiration policies ignore client  timezones and expire globally at the specified expires_at time.", "example": "2014-01-01T05:20:20Z", "type": "string", "x-auditable": true}]}}, "nullable": true, "required": ["expires_at"], "x-stainless-terraform-configurability": "computed_optional"}
```
