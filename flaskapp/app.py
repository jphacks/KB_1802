from flask import Flask
from flaskext.sass import sass

app = Flask(__name__)
sass(app, input_dir='assets/', output_dir='static/css')

@app.route('/')
def hello_world():
	return 'yeew'

if __name__ == "__main__":
    app.run(debug=True)
