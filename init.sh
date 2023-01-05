#!/bin/bash

curl -H "Content-type:application/json" $(gp url 3000)/api/v1/users.register -d '{ "username": "user0", "email": "a@b.com", "pass": "123456", "name": "user"}'