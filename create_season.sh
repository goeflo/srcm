#!/bin/bash

curl -X POST http://localhost:8081/event/season -H 'Content-Type: application/json' -d '{"name":"some season"}'