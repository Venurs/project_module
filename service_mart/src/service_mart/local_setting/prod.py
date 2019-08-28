import os


# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = False

# Database
# https://docs.djangoproject.com/en/2.1/ref/settings/#databases
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'HOST': '',
        'PORT': '3306',
        'NAME': '',
        'USER': 'mart',
        'PASSWORD': '',
        'TEST': {
            'NAME': '',
        },
    },
    'kjzd': {
        'ENGINE': 'django.db.backends.mysql',
        'HOST': '',
        'PORT': '',
        'NAME': '',
        'USER': '',
        'PASSWORD': ''
    }
}

CACHES = {
    "default": {
        "BACKEND": "django_redis.cache.RedisCache",
        "LOCATION": "redis://redis:16379",
        # "LOCATION": "redis://:{user}@{host}:{port}",
        "OPTIONS": {
            "CLIENT_CLASS": "django_redis.client.DefaultClient",
            "PASSWORD": "dWRd4XgT75QtyTLx"
        }
    }
}

MEDIA_ROOT = ''
