#!/bin/bash
echo "Stopping and then deleting existing DB container"
docker stop db
docker container rm -f db

echo "Build data generator"
cd support
g++ generateDataFiles.cpp BinaryTree.cpp -o generate
echo "Generate Data Files.... this may take some time...."
./generate
cd ../

echo "Temporary Data Cleanup Step to handle purchase records with 3 fields"
mkdir -p db/init_data
cd db/init_data
grep -v '^.* .* .*$' purchaseHistory.txt > temp
mv temp purchaseHistory.txt
cd ../../



echo "Clearing out any existing DB data files"
sudo rm -rf /tmp/data
mkdir /tmp/data

echo "Starting mysql DB on port 32000.  Initializing the DB with ./db/init.sql"
docker run --name db -p 32000:3306 -v /tmp/data:/var/lib/mysql \
  -v ./db/init.sql:/docker-entrypoint-initdb.d/init.sql \
  -v ./db/init_data:/docker-entrypoint-initdb.d/init_data \
  -e MYSQL_ROOT_PASSWORD=root -d docker.io/mysql:8.0.22 \
  --secure-file-priv=/docker-entrypoint-initdb.d/init_data

echo "You can connect to this db with the command: mysql -u root --host=localhost --port=32000 --protocol=tcp -p"
