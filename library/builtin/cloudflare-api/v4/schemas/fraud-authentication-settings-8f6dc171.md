---
title: fraud_authentication_settings
page_id: schema-fraud-authentication-settings-8f6dc171
path: schemas
description: |-
    Configuration for classifying login authentication outcomes based on the origin response.
    Requires `user_profiles` to be enabled.

    - Success and failure criteria are independently updatable — sending only `success_criteria`
      leaves failure codes untouched, and vice versa.
    - Omit `authentication_settings` entirely to leave both unchanged.
    - Status codes must not overlap between success and failure criteria.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# fraud_authentication_settings

Configuration for classifying login authentication outcomes based on the origin response.
Requires `user_profiles` to be enabled.

- Success and failure criteria are independently updatable — sending only `success_criteria`
  leaves failure codes untouched, and vice versa.
- Omit `authentication_settings` entirely to leave both unchanged.
- Status codes must not overlap between success and failure criteria.

```yaml
{"description": "Configuration for classifying login authentication outcomes based on the origin response.\nRequires `user_profiles` to be enabled.\n\n- Success and failure criteria are independently updatable — sending only `success_criteria`\n  leaves failure codes untouched, and vice versa.\n- Omit `authentication_settings` entirely to leave both unchanged.\n- Status codes must not overlap between success and failure criteria.\n", "type": "object", "properties": {"failure_criteria": {"description": "Criterion for identifying failed login responses.", "allOf": [{"$ref": "#/components/schemas/fraud_auth_criteria"}]}, "success_criteria": {"description": "Criterion for identifying successful login responses.", "allOf": [{"$ref": "#/components/schemas/fraud_auth_criteria"}]}}, "additionalProperties": false}
```
