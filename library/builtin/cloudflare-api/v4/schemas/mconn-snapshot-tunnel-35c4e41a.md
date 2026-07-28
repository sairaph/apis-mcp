---
title: mconn_snapshot_tunnel
page_id: schema-mconn-snapshot-tunnel-35c4e41a
path: schemas
description: Snapshot Tunnels
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_snapshot_tunnel

Snapshot Tunnels

```yaml
{"description": "Snapshot Tunnels", "type": "object", "properties": {"health_state": {"description": "Name of tunnel health state (unknown, healthy, degraded, down)", "type": "string"}, "health_value": {"description": "Numeric value associated with tunnel state (0 = unknown, 1 = healthy, 2 = degraded, 3 = down)", "type": "number"}, "interface_name": {"description": "The tunnel interface name (i.e. xfrm1, xfrm3.99, etc.)", "type": "string"}, "natd_result": {"description": "Public socket address returned by the NAT detector", "type": "string"}, "natd_state": {"description": "Numeric NAT detector state (0 = detected, 1 = missing result, 2 = stale result)", "type": "number"}, "natd_target": {"description": "Target socket address probed by the NAT detector, using the detector source port", "type": "string"}, "probed_mtu": {"description": "MTU as measured between the two ends of the tunnel", "type": "number"}, "recent_healthy_pings": {"description": "Number of recent healthy pings for this tunnel", "type": "number"}, "recent_unhealthy_pings": {"description": "Number of recent unhealthy pings for this tunnel", "type": "number"}, "tunnel_id": {"description": "Tunnel identifier", "type": "string"}}, "required": ["tunnel_id", "interface_name", "health_state", "health_value"]}
```
