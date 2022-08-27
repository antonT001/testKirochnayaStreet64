#!/bin/bash

ab -n 10000 -c 300 -k -H "Accept-Encoding: gzip, deflate" -p ab_file.data -T application/x-www-form-urlencoded http://127.0.0.1:8000/log