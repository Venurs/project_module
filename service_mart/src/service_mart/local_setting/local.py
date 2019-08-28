import os

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True

# Database
# https://docs.djangoproject.com/en/2.1/ref/settings/#databases
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'HOST': '',
        'PORT': '3306',
        'NAME': '',
        'USER': '',
        'PASSWORD': '',
        'TEST': {
            'NAME': 'test_martservice',
        },
    },
    'kjzd': {
        'ENGINE': 'django.db.backends.mysql',
        'HOST': '',
        'PORT': '3306',
        'NAME': '',
        'USER': '',
        'PASSWORD': ''
    }
}

CACHES = {
    "default": {
        "BACKEND": "django_redis.cache.RedisCache",
        "LOCATION": "redis://redis:6379",
        # "LOCATION": "redis://:{user}@{host}:{port}",
        "OPTIONS": {
            "CLIENT_CLASS": "django_redis.client.DefaultClient",
            # "PASSWORD": "sg123456"
        }
    }
}

MEDIA_ROOT = '/django_file/'

