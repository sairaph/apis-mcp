---
title: mconn_interrupt
page_id: schema-mconn-interrupt-bedbede4
path: schemas
description: Interrupt action for a connector.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_interrupt

Interrupt action for a connector.

```yaml
{"description": "Interrupt action for a connector.", "type": "object", "properties": {"reboot": {"type": "object", "additionalProperties": false, "properties": {"purge": {"description": "Purge connector state.", "type": "boolean", "default": false}}}, "restart": {"type": "object", "additionalProperties": false, "properties": {"purge": {"description": "Purge connector state.", "type": "boolean", "default": false}}}, "shutdown": {"type": "object", "additionalProperties": false, "properties": {"purge": {"description": "Purge connector state.", "type": "boolean", "default": false}}}}, "additionalProperties": false, "maxProperties": 1, "minProperties": 1}
```
