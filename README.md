# nats

```bash 
cd ex1
```

start with one subscriber, one queuesubscriber and publisher
```bash
docker-compose up -d --build
```

scale subscribers to 3
```bash
docker-compose up -d --scale --scale sub_order=3
```

logs for all containers `sub_order`  
```bash
docker-compose logs -f sub_order
```

all services get message

```
sub_order_2          | 2021/09/09 11:21:22 get data  "Hello World" 
sub_order_1          | 2021/09/09 11:21:22 get data  "Hello World" 
sub_order_3          | 2021/09/09 11:21:22 get data  "Hello World" 
```


scale queue subscribers to 3
```
docker-compose up -d --scale queue_sub_order=3
```

logs for all containers `queue_sub_order`

```bash
docker-compose logs -f queue_sub_order
```

only one instance get message  

```
queue_sub_order_2    | 2021/09/09 11:23:36 get data  "Hello World" 
queue_sub_order_1    | 2021/09/09 11:23:38 get data  "Hello World" 
queue_sub_order_3    | 2021/09/09 11:23:40 get data  "Hello World"
```
