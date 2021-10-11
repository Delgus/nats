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

### At least once delivery?

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

