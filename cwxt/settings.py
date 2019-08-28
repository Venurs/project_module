import os
TORNADO_LOCAL_SETTING = os.environ.get('TORNADO_LOCAL_SETTING', 'local')
exec(f'from local_settings.{TORNADO_LOCAL_SETTING} import *')

# There are just a few companies that don't bother to create databases tables

COMPANY_MASTER = {
    1: "深圳市前海必胜道",
    2: "卖家成长",
    3: "亚马逊信息科技"
}