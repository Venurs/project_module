import os





CACHES = {
    'HOST': '',
    'PORT': 6379,
    'PASSWORD': "",
}

BASE_DIR = os.path.dirname(__file__)
settings_dict = dict(
    debug=True,
    template_path=os.path.join(BASE_DIR, "templates"),
    cookie_secret="s5w4ertygyujoikoih7ytf5rtuygy7",
    xsrf_cookies=False,
    autoreload=True
)
