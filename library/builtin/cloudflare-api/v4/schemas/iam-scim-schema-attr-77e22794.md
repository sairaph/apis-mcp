---
title: iam_scim_schema_attr
page_id: schema-iam-scim-schema-attr-77e22794
path: schemas
description: An attribute definition within a SCIM schema (RFC 7643 Section 7).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_schema_attr

An attribute definition within a SCIM schema (RFC 7643 Section 7).

```yaml
{"description": "An attribute definition within a SCIM schema (RFC 7643 Section 7).", "type": "object", "properties": {"canonicalValues": {"description": "A collection of canonical values for the attribute.", "type": "array", "items": {"type": "string"}}, "caseExact": {"description": "Indicates if the string attribute is case-sensitive.", "type": "boolean", "example": false}, "description": {"description": "A human-readable description of the attribute.", "type": "string"}, "multiValued": {"description": "Indicates if the attribute is multi-valued.", "type": "boolean", "example": false}, "mutability": {"description": "Indicates the circumstances under which the value of the attribute can be defined or redefined.", "type": "string", "example": "readWrite", "enum": ["readOnly", "readWrite", "immutable", "writeOnly"]}, "name": {"description": "The attribute's name.", "type": "string", "example": "userName"}, "referenceTypes": {"description": "A multi-valued attribute that indicates the SCIM resource types that may be referenced.", "type": "array", "items": {"type": "string"}}, "required": {"description": "Indicates if the attribute is required.", "type": "boolean", "example": false}, "returned": {"description": "Indicates when an attribute and associated values are returned in response to a GET request or in response to a PUT, POST, or PATCH request.", "type": "string", "example": "default", "enum": ["always", "never", "default", "request"]}, "subAttributes": {"description": "Defines a set of sub-attributes when the attribute type is `complex`.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scim_schema_attr"}}, "type": {"description": "The attribute's data type.", "type": "string", "example": "string", "enum": ["string", "boolean", "decimal", "integer", "dateTime", "reference", "complex"]}, "uniqueness": {"description": "Indicates how the service provider enforces uniqueness of attribute values.", "type": "string", "example": "none", "enum": ["none", "server", "global"]}}, "required": ["name", "type", "multiValued", "description", "required", "caseExact", "mutability", "returned", "uniqueness"], "title": "SCIM Schema Attribute"}
```
