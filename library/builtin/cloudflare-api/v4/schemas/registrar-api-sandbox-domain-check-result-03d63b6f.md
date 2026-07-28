---
title: registrar-api-sandbox_domain_check_result
page_id: schema-registrar-api-sandbox-domain-check-result-03d63b6f
path: schemas
description: Represents a single authoritative domain availability result returned by the Check endpoint. Check results reflect current registry status and should be used immediately before registration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_domain_check_result

Represents a single authoritative domain availability result returned by the Check endpoint. Check results reflect current registry status and should be used immediately before registration.

```yaml
{"description": "Represents a single authoritative domain availability result returned by the Check endpoint. Check results reflect current registry status and should be used immediately before registration.", "type": "object", "properties": {"name": {"description": "The fully qualified domain name (FQDN) in punycode format for internationalized domain names (IDNs).", "type": "string", "example": "example.com"}, "pricing": {"$ref": "#/components/schemas/registrar-api-sandbox_pricing"}, "reason": {"description": "Present only when `registrable` is `false`. Explains why the domain cannot be registered via this API.\n- `extension_not_supported_via_api`: Cloudflare Registrar supports this extension in the dashboard but it is not yet available for programmatic registration via this API. The user can register via `https://dash.cloudflare.com/{account_id}/domains/registrations`.\n- `extension_not_supported`: This extension is not supported by Cloudflare Registrar at all.\n- `extension_disallows_registration`: The extension's registry has temporarily or permanently frozen new registrations. No registrar can register domains on this extension at this time.\n- `domain_premium`: The domain is premium priced. Premium registration is not currently supported by this API.\n- `domain_unavailable`: The domain is already registered, reserved, or otherwise not available on a supported extension.", "type": "string", "example": "domain_unavailable", "enum": ["extension_not_supported_via_api", "extension_not_supported", "extension_disallows_registration", "domain_premium", "domain_unavailable"]}, "registrable": {"description": "Indicates whether this domain can be registered programmatically through this API based on a real-time registry check.\n- `true`: Domain is available for registration. The `pricing` object will be included.\n- `false`: Domain is not available. See the `reason` field for why. `tier` may still be present on some non-registrable results, such as premium domains.", "type": "boolean", "example": true}, "tier": {"description": "The pricing tier for this domain. Always present when `registrable` is `true`; defaults to `standard` for most domains. May be absent when `registrable` is `false`.\n- `standard`: Standard registry pricing\n- `premium`: Premium domain with higher pricing set by the registry", "type": "string", "example": "standard", "enum": ["standard", "premium"]}}, "required": ["name", "registrable"]}
```
