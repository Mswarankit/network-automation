def flatten_json(value):
    flatten = {}
    def flatten_j(x, v = ''):
        if isinstance(x, dict):
            for a in x:
                flatten_j(x[a], v + a + '_')
        else:
            flatten[v[:-1]] = x
    
    flatten_j(value)
    return flatten

def process_flatten_request(data):
    flattened_result = flatten_json(data)
    return flattened_result