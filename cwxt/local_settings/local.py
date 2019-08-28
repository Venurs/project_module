import os

# Give the wrong database
# database = "kjzd_crm"
# host = "47.112.126.125"
# user = "kjzd"
# password = "Kjzd2019!@#!"
# port = 3306
database = "kjzd_oa_dev"
host = "119.23.227.13"
user = "crm_ikjzd_com"
password = "Hcl321123....."
port = 3306

# 供应商系统测试地址
supplier_host = 'http://112.74.170.16:8091'

CACHES = {
    'HOST': 'localhost',
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