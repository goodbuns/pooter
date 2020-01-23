# pooter

Pooter is a simple Twitter clone.

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
curl localhost:8000/users.create --data '{"username":"ba", "password":"hello again"}'
```

Test following a user (make sure 2 users have been created):

```
curl localhost:8000/users.follow --data '{"username":"bab", "idol":"ba"}'
```

Test post creation by user:

```
curl localhost:8000/poots.post --data '{"username":"bab", "password":"hello", "content":"im a sleepo beepo"}'
```

List all posts by a user:

```
curl localhost:8000/users.posts --data '{"username":"bab"}'
```

Check feed of a user (make sure user is following another user with posts). Set `before_time` to some Unix UTC timestamp after posts were created by the users that the user is following.:

```
curl localhost:8000/poots.feed --data '{"username":"bab", "password":"hello", "page_size": 10, "before_time":1578282686}'
```
