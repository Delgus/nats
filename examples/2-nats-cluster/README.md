# Example 2 (cluster)

First peer known about second
```
routes = [
    "nats-route://cluster:clusterpsw@nats-2:6222"
]
```

Second peer known about third.
```
routes = [
    "nats-route://cluster:clusterpsw@nats-3:6222"
]
```

Publisher connect to first peer.

Subscriber connect to third peer.

Start cluster and apps

```bash
docker-compose up -d --build
```

logs for all containers `sub_order`

```bash
docker-compose logs -f sub_order
```

Publish message

```bash
docker-compose exec pub_order ./app
```

Service get message

