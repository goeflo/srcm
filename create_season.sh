#!/bin/bash

curl -X POST http://localhost:8081/event/season -H 'Content-Type: application/json' -d '{"name":"new season"}'
curl -X POST http://localhost:8081/event/season/1/race -H 'Content-Type: application/json' -d '{"name":"new race"}'

curl --data "/home/florian/work/srcm/race01-results.csv" http://localhost:8081/event/race/1/results

 curl -X POST -d '{"email":"admin","passwd":"1234"}' localhost:8080/api/v1/event/season -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluIn0.A9dz8H4vRCdMb39m6nOlnl_HbF5zgof5LrLm2i0xEY0"