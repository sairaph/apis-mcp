---
title: iam_scim_service_provider_config
page_id: schema-iam-scim-service-provider-config-5d045469
path: schemas
description: The SCIM 2.0 Service Provider configuration (RFC 7643 Section 5). Describes which optional SCIM features Cloudflare supports. IdPs use this to auto-configure their SCIM integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_service_provider_config

The SCIM 2.0 Service Provider configuration (RFC 7643 Section 5). Describes which optional SCIM features Cloudflare supports. IdPs use this to auto-configure their SCIM integration.

```yaml
{"description": "The SCIM 2.0 Service Provider configuration (RFC 7643 Section 5). Describes which optional SCIM features Cloudflare supports. IdPs use this to auto-configure their SCIM integration.\n", "type": "object", "properties": {"authenticationSchemes": {"type": "array", "items": {"$ref": "#/components/schemas/iam_scim_authentication_scheme"}}, "bulk": {"$ref": "#/components/schemas/iam_scim_bulk_feature"}, "changePassword": {"$ref": "#/components/schemas/iam_scim_feature"}, "documentationUri": {"description": "An HTTP-addressable URL pointing to the service provider's human-consumable help documentation.", "type": "string", "example": "https://developers.cloudflare.com/fundamentals/account/account-security/scim-setup/"}, "etag": {"$ref": "#/components/schemas/iam_scim_feature"}, "filter": {"$ref": "#/components/schemas/iam_scim_filter_feature"}, "patch": {"$ref": "#/components/schemas/iam_scim_feature"}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"]}, "sort": {"$ref": "#/components/schemas/iam_scim_feature"}}, "required": ["schemas", "patch", "bulk", "filter", "changePassword", "sort", "etag", "authenticationSchemes"], "title": "SCIM Service Provider Config"}
```
