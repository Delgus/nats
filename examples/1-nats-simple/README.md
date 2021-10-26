# Example

```bash
docker-compose up -d --scale sub_order=3 --scale queue_sub_order=3 --build
```

### Subscribe

logs for all containers `sub_order`  

```bash
docker-compose logs -f sub_order
```

all services get message

```
sub_order_1        | 2021/10/26 07:21:12 get data  "nats_message_15" 
sub_order_3        | 2021/10/26 07:21:12 get data  "nats_message_15" 
sub_order_2        | 2021/10/26 07:21:12 get data  "nats_message_15"  
```

### QueueSubscribe

logs for all containers `queue_sub_order`

```bash
docker-compose logs -f queue_sub_order
```

only one instance get message

### At most once delivery

stopped service

```bash
docker-compose stop sub_order
```

started service

```bash
docker-compose start sub_order
```

message lost

```bash
docker-compose logs -f sub_order
```



