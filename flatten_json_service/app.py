from flask import Flask, jsonify, request
from redis import Redis
from rq import Queue
import json
from rq.job import Job
from tasks import process_flatten_request

app = Flask(__name__)
redis_conn = Redis(host='localhost', port=6379)
queue = Queue(connection=redis_conn)



@app.route('/input_request', methods=['POST'])
def submit_flatten_request():
    data = request.json
    job = queue.enqueue(process_flatten_request, data)
    return jsonify({'request_id': job.id}), 202

@app.route('/status/<request_id>', methods=['GET'])
def get_status(request_id):
    try:
        job = Job.fetch(request_id, connection=redis_conn)
        if job.is_finished:
            return jsonify({
                'status': 'completed',
                'result': job.result
            })
        elif job.is_failed:
            return jsonify({
                'status': 'failed',
                'error': str(job.exc_info)
            }), 500
        else:
            return jsonify({
                'status': 'processing'
            }), 202
    except Exception as e:
        return jsonify({
            'status': 'error',
            'message': str(e)
        }), 404

if __name__ == '__main__':
    app.run(debug=True)