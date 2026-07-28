---
title: Create a Customer tax ID
page_id: operation-post-v1-customers-customer-tax-ids-327fdd85
path: operations/untagged
description: <p>Creates a new <code>tax_id</code> object for a customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/customers/{customer}/tax_ids
operation_ids:
    - PostCustomersCustomerTaxIds
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Customer tax ID

`POST /v1/customers/{customer}/tax_ids`

Operation ID: `PostCustomersCustomerTaxIds`

<p>Creates a new <code>tax_id</code> object for a customer.</p>

## Definition

```yaml
{"summary": "Create a Customer tax ID", "description": "<p>Creates a new <code>tax_id</code> object for a customer.</p>", "operationId": "PostCustomersCustomerTaxIds", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["type", "value"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "type": {"maxLength": 5000, "type": "string", "description": "Type of the tax ID, one of `ad_nrt`, `ae_trn`, `al_tin`, `am_tin`, `ao_tin`, `ar_cuit`, `au_abn`, `au_arn`, `aw_tin`, `az_tin`, `ba_tin`, `bb_tin`, `bd_bin`, `bf_ifu`, `bg_uic`, `bh_vat`, `bj_ifu`, `bo_tin`, `br_cnpj`, `br_cpf`, `bs_tin`, `by_tin`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `cd_nif`, `ch_uid`, `ch_vat`, `cl_tin`, `cm_niu`, `cn_tin`, `co_nit`, `cr_tin`, `cv_nif`, `de_stn`, `do_rcn`, `ec_ruc`, `eg_tin`, `es_cif`, `et_tin`, `eu_oss_vat`, `eu_vat`, `fo_vat`, `gb_vat`, `ge_vat`, `gi_tin`, `gn_nif`, `hk_br`, `hr_oib`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `it_cf`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kg_tin`, `kh_tin`, `kr_brn`, `kz_bin`, `la_tin`, `li_uid`, `li_vat`, `lk_vat`, `ma_vat`, `md_vat`, `me_pib`, `mk_vat`, `mr_nif`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `ng_tin`, `no_vat`, `no_voec`, `np_pan`, `nz_gst`, `om_vat`, `pe_ruc`, `ph_tin`, `pl_nip`, `py_ruc`, `ro_tin`, `rs_pib`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `sn_ninea`, `sr_fin`, `sv_nit`, `th_vat`, `tj_tin`, `tr_tin`, `tw_vat`, `tz_vat`, `ua_vat`, `ug_tin`, `us_ein`, `uy_ruc`, `uz_tin`, `uz_vat`, `ve_rif`, `vn_tin`, `za_vat`, `zm_tin`, or `zw_tin`", "enum": ["ad_nrt", "ae_trn", "al_tin", "am_tin", "ao_tin", "ar_cuit", "au_abn", "au_arn", "aw_tin", "az_tin", "ba_tin", "bb_tin", "bd_bin", "bf_ifu", "bg_uic", "bh_vat", "bj_ifu", "bo_tin", "br_cnpj", "br_cpf", "bs_tin", "by_tin", "ca_bn", "ca_gst_hst", "ca_pst_bc", "ca_pst_mb", "ca_pst_sk", "ca_qst", "cd_nif", "ch_uid", "ch_vat", "cl_tin", "cm_niu", "cn_tin", "co_nit", "cr_tin", "cv_nif", "de_stn", "do_rcn", "ec_ruc", "eg_tin", "es_cif", "et_tin", "eu_oss_vat", "eu_vat", "fo_vat", "gb_vat", "ge_vat", "gi_tin", "gn_nif", "hk_br", "hr_oib", "hu_tin", "id_npwp", "il_vat", "in_gst", "is_vat", "it_cf", "jp_cn", "jp_rn", "jp_trn", "ke_pin", "kg_tin", "kh_tin", "kr_brn", "kz_bin", "la_tin", "li_uid", "li_vat", "lk_vat", "ma_vat", "md_vat", "me_pib", "mk_vat", "mr_nif", "mx_rfc", "my_frp", "my_itn", "my_sst", "ng_tin", "no_vat", "no_voec", "np_pan", "nz_gst", "om_vat", "pe_ruc", "ph_tin", "pl_nip", "py_ruc", "ro_tin", "rs_pib", "ru_inn", "ru_kpp", "sa_vat", "sg_gst", "sg_uen", "si_tin", "sn_ninea", "sr_fin", "sv_nit", "th_vat", "tj_tin", "tr_tin", "tw_vat", "tz_vat", "ua_vat", "ug_tin", "us_ein", "uy_ruc", "uz_tin", "uz_vat", "ve_rif", "vn_tin", "za_vat", "zm_tin", "zw_tin"], "x-stripeBypassValidation": true}, "value": {"type": "string", "description": "Value of the tax ID."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax_id"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
