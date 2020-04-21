#!/bin/sh

ab -n 10000 -c 1000 -l http://127.0.0.1:8080/request > request.txt &
ab -n 10000 -c 1000 -l http://127.0.0.1:8080/admin/requests > admin_requests.txt &

