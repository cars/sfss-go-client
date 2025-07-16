#!/bin/bash

export api_gen_version="latest"
export api_gen_language="go"
export sfssapp_spec="SFSSApp_APIs-1.4.0-fixed.json"
export restapi_spec="Final_RestService_JSON-1.4.0-fixed.json"
export sfssapp_pkg_name="sfssapp"
export sfssrest_pkg_name="sfssrest"

docker run --rm -v "$(pwd):/local" \
  openapitools/openapi-generator-cli:${api_gen_version} \
  generate \
  -g ${api_gen_language} \
  -i /local/${sfssapp_spec} \
  -o /local/${sfssapp_pkg_name} \
  --package-name ${sfssapp_pkg_name} \

docker run --rm -v "$(pwd):/local" \
  openapitools/openapi-generator-cli:${api_gen_version} \
  chown -R $UID /local/${sfssapp_pkg_name}


docker run --rm -v "$(pwd):/local" \
  openapitools/openapi-generator-cli:${api_gen_version} \
  generate \
  -g ${api_gen_language} \
  -i /local/${restapi_spec} \
  -o /local/${sfssrest_pkg_name} \
  --package-name ${sfssrest_pkg_name} \

docker run --rm -v "$(pwd):/local" \
  openapitools/openapi-generator-cli:${api_gen_version} \
  chown -R $UID /local/${sfssrest_pkg_name}

  # Run hacks if appropriate

  