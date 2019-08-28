import os

database = "kjzd_oa_dev"
host = "119.23.227.13"
user = "crm_ikjzd_com"
password = "Hcl321123....."
port = 3306

# 供应商系统测试地址
supplier_host = 'http://112.74.170.16:8091'

CACHES = {
    'HOST': '192.168.0.31',
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
