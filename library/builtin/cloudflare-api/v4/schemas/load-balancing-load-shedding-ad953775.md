---
title: load-balancing_load_shedding
page_id: schema-load-balancing-load-shedding-ad953775
path: schemas
description: Configures load shedding policies and percentages for the pool.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_load_shedding

Configures load shedding policies and percentages for the pool.

```yaml
{"description": "Configures load shedding policies and percentages for the pool.", "type": "object", "properties": {"default_percent": {"description": "The percent of traffic to shed from the pool, according to the default policy. Applies to new sessions and traffic without session affinity.", "type": "number", "default": 0, "maximum": 100, "minimum": 0, "x-auditable": true}, "default_policy": {"description": "The default policy to use when load shedding. A random policy randomly sheds a given percent of requests. A hash policy computes a hash over the CF-Connecting-IP address and sheds all requests originating from a percent of IPs.", "type": "string", "default": "random", "enum": ["random", "hash"], "x-auditable": true}, "session_percent": {"description": "The percent of existing sessions to shed from the pool, according to the session policy.", "type": "number", "default": 0, "maximum": 100, "minimum": 0, "x-auditable": true}, "session_policy": {"description": "Only the hash policy is supported for existing sessions (to avoid exponential decay).", "type": "string", "default": "hash", "enum": ["hash"], "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "computed_optional"}
```
