---
title: load-balancing_origin_steering
page_id: schema-load-balancing-origin-steering-fde8444e
path: schemas
description: Configures origin steering for the pool. Controls how origins are selected for new sessions and traffic without session affinity.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_origin_steering

Configures origin steering for the pool. Controls how origins are selected for new sessions and traffic without session affinity.

```yaml
{"description": "Configures origin steering for the pool. Controls how origins are selected for new sessions and traffic without session affinity.", "type": "object", "properties": {"policy": {"description": "The type of origin steering policy to use.\n- `\"random\"`: Select an origin randomly.\n- `\"hash\"`: Select an origin by computing a hash over the CF-Connecting-IP address.\n- `\"least_outstanding_requests\"`: Select an origin by taking into consideration origin weights, as well as each origin's number of outstanding requests. Origins with more pending requests are weighted proportionately less relative to others.\n- `\"least_connections\"`: Select an origin by taking into consideration origin weights, as well as each origin's number of open connections. Origins with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.", "type": "string", "default": "random", "enum": ["random", "hash", "least_outstanding_requests", "least_connections"], "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "computed_optional"}
```
