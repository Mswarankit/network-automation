#### Handling Flatten Json using Redis, Redis Queue and Flask for api endpoints


1. We need to have virtualenv to work on the IDE - (VSCode)

`virtualenv fenv`
this will create an virtual environment for the project.

we activate using `source ./fenv/bin/activate` --for linux users

Install all the `requirements.txt` which contains all the importent package to run the application

```
pip install -r requirements.txt
```

we have install redis server in our system if we don't have one
```
sudo pacman -S redis

sudo systemctl start redis

redis-cli ping
```
`PONG` if we got response our redis server is working.

2. There is app.py and worker.py which are going to run simultaneously to get output

To run in one terminal (we can use tmux for multiple sessions)
``` 
app.py

python  app.py
```
To handle worker side
```
worker.py

python worker.py
```

3. To test our service \
**Request**
```
curl --location 'http://localhost:5000/input_request' \
--header 'Content-Type: application/json' \
--data '{
    "official": {
            "name": {
                    "first": "ankit",
                    "middle": "k",
                    "last": "mishra"
                }
            }
        }'
```
**Response**
```
{
    "request_id": "4c8f4ec1-99d1-4938-a3fe-4440cd4b4243"
}
```

**Request** for getting status of our input request with request_id

```
curl --location 'http://localhost:5000/status/4c8f4ec1-99d1-4938-a3fe-4440cd4b4243'
```

**Response** 
```
{
    "result": {
        "official_name_first": "ankit",
        "official_name_last": "mishra",
        "official_name_middle": "k"
    },
    "status": "completed"
}
```
