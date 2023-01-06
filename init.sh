#!/bin/bash

registerUser() {
    result=$(curl -sS -f -H "Content-type:application/json" http://localhost:3000/api/v1/users.register -d '{ "username": "user1", "email": "b@b.com", "pass": "123456", "name": "user2"}') | jq -r '.success'
    result=true
    if [ "$result" = "true" ] 
    then
        return 1
    else
        return -1
    fi
}

echo "Shell Script Initiated"

success=false
iterations=0
increase=1

while [ "$success" != "1" ] && [ $iterations -lt 20 ]
do
    sleep 1
    echo $(registerUser)
    
    echo $success
    iterations=$(( $iterations+1 ))
done

gp sync-done user-init