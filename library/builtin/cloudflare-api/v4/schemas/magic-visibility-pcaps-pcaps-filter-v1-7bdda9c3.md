---
title: magic-visibility-pcaps_pcaps_filter_v1
page_id: schema-magic-visibility-pcaps-pcaps-filter-v1-7bdda9c3
path: schemas
description: The packet capture filter. When this field is empty, all packets are captured.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-pcaps_pcaps_filter_v1

The packet capture filter. When this field is empty, all packets are captured.

```yaml
{"description": "The packet capture filter. When this field is empty, all packets are captured.", "type": "object", "properties": {"destination_address": {"description": "The destination IP address of the packet.", "type": "string", "example": "1.2.3.4", "x-auditable": true}, "destination_port": {"description": "The destination port of the packet.", "type": "number", "example": 80, "x-auditable": true}, "protocol": {"description": "The protocol number of the packet.", "type": "number", "example": 6, "x-auditable": true}, "source_address": {"description": "The source IP address of the packet.", "type": "string", "example": "1.2.3.4", "x-auditable": true}, "source_port": {"description": "The source port of the packet.", "type": "number", "example": 123, "x-auditable": true}}}
```
