# -*- coding=utf-8 -*-
import pickle

class CacheBackend(object):
    """ The base class of Cache Backend"""

    def get(self, key):
        raise NotImplementedError

    def set(self, key, value, timeout):
        raise NotImplementedError

    def delete(self, key):
        raise NotImplementedError

    def exists(self, key):
        raise NotImplementedError


class RedisCacheBackend(CacheBackend):
    """ The class of Redis Cache Backend"""

    def __init__(self, redis_connection, **kwargs):
        self.kwargs = dict(timeout=24 * 60 * 60)
        self.kwargs.update(kwargs)
        self.redis = redis_connection

    def get(self, key, default=None):
        if self.exists(key):
            return pickle.loads(self.redis.get(key))
        return None

    def set(self, key, value, timeout=None):
        self.redis.set(key, pickle.dumps(value))
        if timeout:
            self.redis.expire(key, timeout)
        else:
            self.redis.expire(key, self.kwargs["timeout"])

    def delete(self, key):
        self.redis.delete(key)

    def exists(self, key):
        return bool(self.redis.exists(key))