from rest_framework.exceptions import ValidationError
from rest_framework.response import Response
from rest_framework import status
from django.http.response import JsonResponse

import os
import traceback
from helper.log import get_logger
from base.exceptions import ValidateException, LogicException

logger = get_logger(__name__)


def common_api(func):
    def _func(*args, **kwargs):
        try:
            resp = func(*args, **kwargs)
            data = {
                "data": resp.data,
                "code": "ok",
                "message": {},
                "version": 2.0
            }
            resp.data = data
            return resp
        except ValidationError as e:
            data = {'data': None, "code": "validate", "message": {}, "version": 2.0}
            if isinstance(e.detail, list):
                data['message']["error_message"] = ' '.join(e.detail)
            else:
                for k, v in e.detail.items():
                    if not isinstance(v, list):
                        v = [v]
                    data['message'][k] = ' '.join(v)
                # data['message'] = e.detail
            return Response(data)
        except ValidateException as e:
            msg = e.msg
            code = e.code
            return Response({'data':    None,
                             "code":    code,
                             "message": msg,
                             "version": 2.0})
        except LogicException as e:
            msg = e.msg
            code = e.code
            return Response({'data':    None,
                             "code":    code,
                             "message": {code: msg},
                             "version": 2.0})
        except Exception as e:
            logger.error('error:error response:{}'.format(e))
            if os.environ.get('DJANGO_LOCAL_SETTING') == 'prod':
                msg = '未知错误，请联系技术人员'
                code = 'error:exception'
                return Response({
                    'data': None,
                    'code': code,
                    'message': {code: msg},
                    'version': 2.0
                })
            else:
                raise
    return _func


def file_response_common_api(func):

    def _func(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except ValidationError as e:
            data = {'data': None, "code": "validate", "message": {}, "version": 2.0}
            for k, v in e.detail.items():
                if not isinstance(v, list):
                    v = [v]
                data['message'][k] = ' '.join(v)
            return Response(data)
        except ValidateException as e:
            msg = e.msg
            code = e.code
            return Response({'data':    None,
                             "code":    code,
                             "message": msg,
                             "version": 2.0}, status=status.HTTP_200_OK)
        except LogicException as e:
            msg = e.msg
            code = e.code
            return Response({'data':    None,
                             "code":    code,
                             "message": {code: msg},
                             "version": 2.0})
        except Exception as e:
            logger.error('error:error response:{}'.format(traceback.format_exc()))
            if os.environ.get('DJANGO_LOCAL_SETTING') == 'prod':
                msg = '未知错误，请联系技术人员'
                code = 'error:exception'
                return Response({
                    'data': None,
                    'code': code,
                    'message': {code: msg},
                    'version': 2.0
                })
            else:
                raise e

    return _func
