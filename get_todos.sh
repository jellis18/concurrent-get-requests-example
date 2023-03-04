#!/bin/bash

TODO_SERVER="https://jsonplaceholder.typicode.com/todos"
NUM_TODOS=$1


get_todo(){
    sleep 0.5
    curl -s --fail $TODO_SERVER/$1 | jq -r '.title' > "/tmp/todo_$1.txt"
}

for ((i=1; i<=${NUM_TODOS}; i++));
do
    get_todo $i &
done

wait

echo "# Todos"
for ((i=1; i<=${NUM_TODOS}; i++));
do
    echo "* $(cat /tmp/todo_${i}.txt)" && rm "/tmp/todo_${i}.txt"
done