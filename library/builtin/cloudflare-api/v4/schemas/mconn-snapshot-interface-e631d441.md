---
title: mconn_snapshot_interface
page_id: schema-mconn-snapshot-interface-e631d441
path: schemas
description: Snapshot Interface
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_snapshot_interface

Snapshot Interface

```yaml
{"description": "Snapshot Interface", "type": "object", "properties": {"ip_addresses": {"type": "array", "items": {"$ref": "#/components/schemas/mconn_snapshot_interface_address"}}, "name": {"description": "Name of the network interface", "type": "string"}, "operstate": {"description": "UP/DOWN state of the network interface", "type": "string"}, "speed": {"description": "Speed of the network interface (bits per second)", "type": "number"}}, "required": ["name", "operstate"]}
```
