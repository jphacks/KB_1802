from flask import Flask, render_template
from flask.ext.assets import Environment, Bundle

app = Flask(__name__)

assets = Environment(app)
assets.url = app.static_url_path
scss = Bundle('soisy', filters='pyscss', output='all.css')
assets.register('scss_all', scss)

@app.route('/', methods=['POST'])
def hello_world():
	return 'yeew'

if __name__ == "__main__":
    app.run(debug=True)
