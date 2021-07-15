DOC_PORT=8890
DOC_PATH=/openapiv2/host.swagger.json
SVC_NAME=host-server
DOC_PREFIX=docs-

docker run \
  -p $DOC_PORT:8080 \
  -v $PWD/openapiv2/:/openapiv2 \
  -e SWAGGER_JSON=$DOC_PATH \
  --name=$DOC_PREFIX$SVC_NAME \
  swaggerapi/swagger-ui
