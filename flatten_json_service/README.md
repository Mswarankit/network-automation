### Handling Flatten Json using Redis, Redis Queue and Flask
This project demonstrates how to process and handle flattened JSON using a Redis-backed queue system and Flask for API endpoints.

---
#### Setup Instruction


##### 1. Create a Virtual Environment 

```
virtualenv fenv
```

Activate the virtual environment (Linux):

```
source ./fenv/bin/activate
```

Install all the `requirements.txt` which contains all the importent package to run the application

```
pip install -r requirements.txt
```

we have install redis server in our system if we don't have one

#### Install Redis

1. Install Redis
```
sudo pacman -S redis
```
2. Start the Redis
```
sudo systemctl start redis
```
3. Verify redis is running
```
redis-cli ping
```
`PONG` if we got response our redis server is working.


#### Running the Application

Running the Application
The project consists of two main components:

`app.py`: The Flask application handles incoming API requests.
`worker.py`: The worker processes tasks from the Redis queue.
Both components need to run simultaneously.

__Running the Flask Application__ \
Run the Flask application in one terminal:
```
python app.py
```

Running the Worker
In a second terminal, start the worker process:
```
python worker.py
```
Note: You can use tools like tmux or screen to manage multiple terminal sessions for convenience.

#### Testing the APIs
__Input Request__ \
Send a POST request to the /input_request endpoint to submit a JSON payload for processing.

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

- The request_id is a unique identifier for tracking the processing status of this input.


__Check Request Status__ \
Use the request_id to check the status and result of the processing by making a GET request to the /status/<request_id> endpoint.

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
