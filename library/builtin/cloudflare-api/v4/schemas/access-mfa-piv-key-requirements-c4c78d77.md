---
title: access_mfa_piv_key_requirements
page_id: schema-access-mfa-piv-key-requirements-c4c78d77
path: schemas
description: Configures PIV key requirements for MFA using hardware security keys.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_mfa_piv_key_requirements

Configures PIV key requirements for MFA using hardware security keys.

```yaml
{"description": "Configures PIV key requirements for MFA using hardware security keys.", "type": "object", "properties": {"pin_policy": {"description": "Defines when a PIN is required to use the SSH key. Valid values: `never` (no PIN required), `once` (PIN required once per session), `always` (PIN required for each use).", "type": "string", "example": "always", "enum": ["never", "once", "always"]}, "require_fips_device": {"description": "Requires the PIV key to be stored on a FIPS 140-2 Level 1 or higher validated device.", "type": "boolean", "example": true}, "ssh_key_size": {"description": "Specifies the allowed SSH key sizes in bits. Valid sizes depend on key type. Ed25519 has a fixed key size and does not accept this parameter.", "type": "array", "items": {"enum": [256, 384, 521, 2048, 3072, 4096], "type": "integer"}, "example": [256, 2048]}, "ssh_key_type": {"description": "Specifies the allowed SSH key types. Valid values are `ecdsa`, `ed25519`, and `rsa`.", "type": "array", "items": {"enum": ["ecdsa", "ed25519", "rsa"], "type": "string"}, "example": ["ecdsa", "rsa"]}, "touch_policy": {"description": "Defines when physical touch is required to use the SSH key. Valid values: `never` (no touch required), `always` (touch required for each use), `cached` (touch cached for 15 seconds).", "type": "string", "example": "always", "enum": ["never", "always", "cached"]}}, "x-auditable": true}
```
