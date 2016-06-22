#!/bin/bash
echo "============ start mongodb ====="

mongod --dbpath $Db_PATH

echo "============ end mongodb ======="

echo "============ start ObjectId ====="

cd $ObjectId_PATH/
bee run

echo "============ end ObjectId ======="
