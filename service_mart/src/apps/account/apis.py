# coding=utf-8

from rest_framework.response import Response
from base.exceptions import LogicException, ValidateException
from django.contrib.auth import logout, login
from django.http import HttpResponseRedirect
from rest_framework.decorators import api_view

from account.models import KjzdUser, MyUser
from account.utils import str2md5
from common.decorator import common_api


@api_view(['POST'])
@common_api
def log_in(request):
    user_data = request.data
    name = user_data.get('username', None)
    pwd = user_data.get('password', None)

    if not all((name, pwd)):
        raise LogicException('error:error', 'Incomplete Params!')
    try:
        kjzd_user = KjzdUser.objects.using('kjzd').get(tel=name)
    except Exception:
        raise LogicException('error:error', 'User Not Found')

    if kjzd_user.password == str2md5(str(pwd)):
        user, _ = MyUser.objects.get_or_create(kjzd_user_id=kjzd_user.id,
                                               defaults={'email': kjzd_user.email,
                                                         'name': kjzd_user.nick,
                                                         'phone': kjzd_user.tel,
                                                         'password':kjzd_user.password})
        login(request, user)
        user_info = {
            'nickname': user.name,
            'phone': user.phone,
            'kjzd_user_id': user.kjzd_user_id,
            'is_manager': False
        }
        if user.has_perm('account.can_login_cms'):
            user_info['is_manager'] = True

        return Response(user_info)
    else:
        raise LogicException('error:error', 'Password Not Right')


def log_out(request):
    logout(request)
    return HttpResponseRedirect('/')