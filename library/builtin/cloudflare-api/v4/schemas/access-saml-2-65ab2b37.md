---
title: access_saml-2
page_id: schema-access-saml-2-65ab2b37
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saml-2

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_identity-provider-2"}, {"properties": {"config": {"type": "object", "properties": {"attributes": {"description": "A list of SAML attribute names that will be added to your signed JWT token and can be used in SAML policy rules.", "type": "array", "items": {"type": "string"}, "example": ["group", "department_code", "divison"]}, "email_attribute_name": {"description": "The attribute name for email in the SAML response.", "type": "string", "example": "Email"}, "header_attributes": {"description": "Add a list of attribute names that will be returned in the response header from the Access callback.", "type": "array", "items": {"properties": {"attribute_name": {"description": "attribute name from the IDP", "type": "string"}, "header_name": {"description": "header that will be added on the request to the origin", "type": "string"}}, "type": "object"}}, "idp_public_certs": {"description": "X509 certificate to verify the signature in the SAML authentication response", "type": "array", "items": {"type": "string"}}, "issuer_url": {"description": "IdP Entity ID or Issuer URL", "type": "string", "example": "https://whoami.com"}, "sign_request": {"description": "Sign the SAML authentication request with Access credentials. To verify the signature, use the public key from the Access certs endpoints.", "type": "boolean"}, "sso_target_url": {"description": "URL to send the SAML authentication requests to", "type": "string", "example": "https://edgeaccess.org/idp/saml/login"}}}}, "type": "object"}], "title": "Generic SAML"}
```
