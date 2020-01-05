# pooter

Pooter is Twitter for thots and poots.

## Development

To run:

```
docker-compose -f ./deployments/pooter-dev/docker-compose.yml up --build
```

To clean:

```
docker container prune && docker volume prune
```

## Testing

Test user creation:

```
curl localhost:8000/users.create --data '{"username":"test2", "password":"hello"}'
```

Test following a user (make sure 2 users have been created):

```
curl localhost:8000/users.follow --data '{"user_id":"1", "follow_id":"2"}'
```
