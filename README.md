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
curl localhost:8000/users.create --data '{"username":"bab", "password":"hello"}'
```

Test following a user (make sure 2 users have been created):

```
curl localhost:8000/users.follow --data '{"user_id":"1", "follow_id":"2"}'
```

Test post creation by user:

```
curl localhost:8000/poots.post --data '{"user_id":"1", "password":"hello", "content":"im a sleepo beepo"}'
```

List all posts by a user:

```
curl localhost:8000/users.posts --data '{"user_id":"1"}'
```

Check feed of a user (make sure user is following another user with posts):

```
curl localhost:8000/poots.feed --data '{"user_id":"1", "password":"hello", "limit": 10, "page":0}'
```
