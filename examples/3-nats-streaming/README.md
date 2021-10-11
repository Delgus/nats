# Example 3 (NATS Streaming)

### Subscribe 

start containers

```bash 
docker-compose up -d --build
```

logs for `sub_order`

```bash
docker-compose logs -f sub_order
```

### At least once delivery

```bash
docker-compose logs -f sub_order
```

```
sub_order_1       | 2021/10/11 08:49:59 get data  "Hello from streaming number:1" 
sub_order_1       | 2021/10/11 08:50:04 get data  "Hello from streaming number:2" 
sub_order_1       | 2021/10/11 08:50:09 get data  "Hello from streaming number:3" 
sub_order_1       | 2021/10/11 08:50:14 get data  "Hello from streaming number:4" 

```

stopped `sub_order`  

```bash
docker-compose stop sub_order
```

started `sub_order`  

```bash
docker-compose start sub_order
```

logs for `sub_order`  

```bash
docker-compose logs -f sub_order  
```

get lost messages

```bash
sub_order_1       | 2021/10/11 08:50:14 get data  "Hello from streaming number:4" 
sub_order_1       | 2021/10/11 08:50:44 get data  "Hello from streaming number:5" 
sub_order_1       | 2021/10/11 08:50:44 get data  "Hello from streaming number:6" 
sub_order_1       | 2021/10/11 08:50:44 get data  "Hello from streaming number:7" 
sub_order_1       | 2021/10/11 08:50:44 get data  "Hello from streaming number:8" 
```
