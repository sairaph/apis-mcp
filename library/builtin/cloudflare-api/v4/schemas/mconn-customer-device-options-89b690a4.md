---
title: mconn_customer_device_options
page_id: schema-mconn-customer-device-options-89b690a4
path: schemas
description: Exactly one of id, serial_number, or provision_license must be provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_customer_device_options

Exactly one of id, serial_number, or provision_license must be provided.

```yaml
{"description": "Exactly one of id, serial_number, or provision_license must be provided.", "type": "object", "properties": {"id": {"type": "string", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "provision_license": {"description": "When true, create and provision a new licence key for the connector.", "type": "boolean"}, "serial_number": {"type": "string", "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "maxProperties": 1, "minProperties": 1}
```
