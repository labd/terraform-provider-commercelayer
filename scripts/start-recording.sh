curl --location --request POST 'http://localhost:8080/__admin/recordings/start' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "targetBaseUrl": "https://labdigital-toyota-poc.commercelayer.io"
}'
