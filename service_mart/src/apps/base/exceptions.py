from django.utils.translation import ugettext_lazy as _
from rest_framework import status


class ParamValueError(Exception):

    def __init__(self, message=_('Param value error.'), code=status.HTTP_400_BAD_REQUEST, params=None):
        self.message = message
        self.code = code
        self.params = params
        super(ParamValueError, self).__init__(message, code, params)


class UnExpectedValueError(Exception):

    def __init__(self, message=_('Unexpected value error.'), code=status.HTTP_400_BAD_REQUEST, params=None):
        self.message = message
        self.code = code
        self.params = params
        super(UnExpectedValueError, self).__init__(message, code, params)


class CustomException(Exception):
    def __init__(self, errmsg, errcode=4000000, *args):
        self.errmsg = errmsg
        self.errcode = errcode
        super(CustomException, self).__init__(errmsg, *args)


class CustomApiException(CustomException):
    def __init__(self, errmsg, errcode=4000000, status_code=200, *args):
        super(CustomApiException, self).__init__(errmsg, errcode, *args)
        self.status_code = status_code


class InvalidApiParamError(CustomApiException):
    def __init__(self, detail, status_code=200):
        super(InvalidApiParamError, self).__init__(detail, 4000111, status_code)


class RPCCallTimeoutException(CustomException):
    pass


class LogicException(Exception):
    """
    逻辑异常类
    """
    def __init__(self, code, msg):
        """
        初始化逻辑异常类
        :param code: 消息码（由异常级别，冒号，异常名组成）
        :param msg: 消息内容
        """
        if isinstance(msg, str) and isinstance(code, str):
            self.msg = msg
            self.code = code
        else:
            self.msg = "消息格式错误"
            self.code = "error:message_code_error"
        super(LogicException, self).__init__(code, msg)


class ValidateException(Exception):
    """
    校验异常类
    """
    def __init__(self, *args, **kwargs):
        """
        初始化校验异常类
        :param args: 无
        :param kwargs: 初始异常信息
        """
        self.msg = kwargs
        self.code = "validate"
        super(ValidateException, self).__init__(*args)

    def add_message(self, key, value):
        """
        添加消息
        :param key:
        :param value:
        """
        self.msg[key] = value
        return self

    def remove_message(self, key):
        self.msg.pop(key, None)
        return self

    def is_error(self):
        return bool(self.msg)

    def add_message_dict(self, message_dict):
        self.msg.update(message_dict)
        return self