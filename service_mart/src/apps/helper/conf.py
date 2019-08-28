from appconf import AppConf
from django.conf import settings
from django.core.exceptions import PermissionDenied
from django.http import Http404


class HelperAppConf(AppConf):

    logging_cookie_domain = None

    logging_cookie_sid_name = 'sid'
    logging_session_sid_name = 'helper.middleware.sid'
    logging_cookie_bid_name = 'bid'
    logging_cookie_kjzd_user_id_name = 'kjzd_user_id'

    logging_except_exceptions = (Http404, SystemExit, PermissionDenied)

    @staticmethod
    def configure_logging_cookie_domain(value):
        return getattr(settings, 'LOGGING_COOKIE_DOMAIN', None)

    class Meta:
        prefix = 'helper'


helper_settings = HelperAppConf()
