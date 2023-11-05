cd controller
swag init
cd ..
cp -r controller/docs .
rm -r controller/docs
docker-compose up