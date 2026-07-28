---
title: registrar-api_domain_search_result
page_id: schema-registrar-api-domain-search-result-dc96a741
path: schemas
description: Represents a single domain suggestion returned by the Search endpoint. Search results are non-authoritative and may be based on cached data. Use POST /domain-check to confirm real-time availability and pricing before registration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_domain_search_result

Represents a single domain suggestion returned by the Search endpoint. Search results are non-authoritative and may be based on cached data. Use POST /domain-check to confirm real-time availability and pricing before registration.

```yaml
{"description": "Represents a single domain suggestion returned by the Search endpoint. Search results are non-authoritative and may be based on cached data. Use POST /domain-check to confirm real-time availability and pricing before registration.", "type": "object", "properties": {"name": {"description": "The fully qualified domain name (FQDN) in punycode format for internationalized domain names (IDNs).", "type": "string", "example": "example.com"}, "pricing": {"$ref": "#/components/schemas/registrar-api_pricing"}, "reason": {"description": "Present only when `registrable` is `false` on search results. Explains why the domain does not appear registrable through this API. These values are advisory; use POST /domain-check for authoritative status.\n- `extension_not_supported_via_api`: Cloudflare Registrar supports this extension in the dashboard but it is not yet available for programmatic registration via this API.\n- `extension_not_supported`: This extension is not supported by Cloudflare Registrar at all.\n- `extension_disallows_registration`: The extension's registry has temporarily or permanently frozen new registrations.\n- `domain_premium`: The domain is premium priced. Premium registration is not currently supported by this API.\n- `domain_unavailable`: The domain appears unavailable.", "type": "string", "example": "domain_unavailable", "enum": ["extension_not_supported_via_api", "extension_not_supported", "extension_disallows_registration", "domain_premium", "domain_unavailable"]}, "registrable": {"description": "Indicates whether this domain appears available based on search data. Search results are non-authoritative and may be stale. - `true`: The domain appears available. Use POST /domain-check to confirm before registration.\n- `false`: The domain does not appear available in search results.", "type": "boolean", "example": true}, "tier": {"description": "The pricing tier for this domain. Always present when `registrable` is `true`;\ndefaults to `standard` for most domains. May be absent when `registrable`\nis `false`.\n- `standard`: Standard registry pricing\n- `premium`: Premium domain with higher pricing set by the registry\n", "type": "string", "example": "standard", "enum": ["standard", "premium"]}}, "required": ["name", "registrable"]}
```
