import os

# Server Socket
bind = '0.0.0.0:8000'
# bind = '0.0.0.0:' + os.getenv('PORT')

# Worker Processes
workers = 2 * os.cpu_count() + 1
threads = 2 * os.cpu_count() + 1
worker_class = 'uvicorn.workers.UvicornWorker'

# Debugging
reload = True

# Logging
accesslog = './logs/access.log'
errorlog = './logs/error.log'
loglevel = 'info'
