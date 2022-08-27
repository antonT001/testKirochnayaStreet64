#!/bin/bash

ab -n 1000 -c 10 -k -H "Accept-Encoding: gzip, deflate" -p ab_file.data -T application/x-www-form-urlencoded http://127.0.0.1:8000/log