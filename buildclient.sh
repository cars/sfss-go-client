#!/bin/bash

export api_gen_version="latest"
export api_gen_language="go"
export openapi_spec="SFSSApp_APIs-1.4.0-fixed.json"
export pkg_name="sfss"

docker run --rm -v "$(pwd):/local" \
  openapitools/openapi-generator-cli:${api_gen_version} \
  generate \
  -g ${api_gen_language} \
  -i /local/${openapi_spec} \
  -o /local/${pkg_name} \
  --package-name ${pkg_name} \

docker run --rm -v "$(pwd):/local" \
  openapitools/openapi-generator-cli:${api_gen_version} \
  chown -R $UID /local/${pkg_name}

  # Run hacks if appropriate

  