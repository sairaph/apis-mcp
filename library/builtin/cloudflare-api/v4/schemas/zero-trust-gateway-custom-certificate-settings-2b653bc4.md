---
title: zero-trust-gateway_custom-certificate-settings
page_id: schema-zero-trust-gateway-custom-certificate-settings-2b653bc4
path: schemas
description: Specify custom certificate settings for BYO-PKI. This field is deprecated; use `certificate` instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_custom-certificate-settings

Specify custom certificate settings for BYO-PKI. This field is deprecated; use `certificate` instead.

```yaml
{"description": "Specify custom certificate settings for BYO-PKI. This field is deprecated; use `certificate` instead.", "type": "object", "properties": {"binding_status": {"description": "Indicate the internal certificate status.", "type": "string", "example": "pending_deployment", "readOnly": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed"}, "enabled": {"description": "Specify whether to enable a custom certificate authority for signing Gateway traffic.", "type": "boolean", "example": true, "nullable": true, "x-auditable": true}, "id": {"description": "Specify the UUID of the certificate (ID from MTLS certificate store).", "type": "string", "example": "d1b364c5-1311-466e-a194-f0e943e0799f", "x-auditable": true}, "updated_at": {"type": "string", "format": "date-time", "readOnly": true, "x-stainless-terraform-configurability": "computed"}}, "deprecated": true, "nullable": true, "required": ["enabled"], "x-stainless-terraform-configurability": "optional"}
```
