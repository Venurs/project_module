import os

database = ""
host = ""
user = ""
password = ""
port = 3306

# 供应商系统测试地址
supplier_host = ''

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
    login_url="/api/log_in/",
    xsrf_cookies=False,
    autoreload=True
)
