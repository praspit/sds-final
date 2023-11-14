# sds-final

## Build docker image of a service

replace shop by the name of your service ( structure need to be the same as this github)
```
docker build -t shop-service --build-arg SERVICE_NAME="shop" -f apps/Dockerfile .
```

then push that image to dockerhub ( preferred name -> sds-{SERVICE_NAME}-service )
