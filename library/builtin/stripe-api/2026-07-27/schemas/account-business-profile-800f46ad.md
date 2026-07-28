---
title: account_business_profile
page_id: schema-account-business-profile-800f46ad
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_business_profile

```yaml
{"title": "AccountBusinessProfile", "type": "object", "properties": {"annual_revenue": {"description": "The applicant's gross annual revenue for its preceding fiscal year.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/account_annual_revenue"}]}, "estimated_worker_count": {"type": "integer", "description": "An estimated upper bound of employees, contractors, vendors, etc. currently working for the business.", "nullable": true}, "mcc": {"maxLength": 5000, "type": "string", "description": "[The merchant category code for the account](/connect/setting-mcc). MCCs are used to classify businesses based on the goods or services they provide.", "nullable": true}, "minority_owned_business_designation": {"type": "array", "description": "Whether the business is a minority-owned, women-owned, and/or LGBTQI+ -owned business.", "nullable": true, "items": {"type": "string", "enum": ["lgbtqi_owned_business", "minority_owned_business", "none_of_these_apply", "prefer_not_to_answer", "women_owned_business"]}}, "monthly_estimated_revenue": {"$ref": "#/components/schemas/account_monthly_estimated_revenue"}, "name": {"maxLength": 5000, "type": "string", "description": "The customer-facing business name.", "nullable": true}, "product_description": {"maxLength": 40000, "type": "string", "description": "Internal-only description of the product sold or service provided by the business. It's used by Stripe for risk and underwriting purposes.", "nullable": true}, "support_address": {"description": "A publicly available mailing address for sending support issues to.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "support_email": {"maxLength": 5000, "type": "string", "description": "A publicly available email address for sending support issues to.", "nullable": true}, "support_phone": {"maxLength": 5000, "type": "string", "description": "A publicly available phone number to call with support issues.", "nullable": true}, "support_url": {"maxLength": 5000, "type": "string", "description": "A publicly available website for handling support issues.", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "The business's publicly available website.", "nullable": true}}, "description": "", "x-expandableFields": ["annual_revenue", "monthly_estimated_revenue", "support_address"]}
```
