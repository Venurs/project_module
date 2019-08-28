import uuid
import sys
import datetime
import json

from django.core.cache import cache
from django.utils import timezone
from django.conf import settings
import requests

from helper.conf import helper_settings
from helper.log import get_logger
from common.mixin import MiddlewareMixin

logger = get_logger(__name__)


class LoggingMiddleware(MiddlewareMixin):
    LOGGER = get_logger('actions')
    RESPONSE_LOG = '%s%s%s' % (
        '{user_id}|{ip}|{bid}|{sid}|{kjzd_user_id}',
        '"{request_method} {request_url}{query_string} {protocol} {status_code} {content_type} {referrer}"',
        '|{ua}'
    )

    def __init__(self, get_response=None):
        super(LoggingMiddleware, self).__init__(get_response)
        self.SGUID_EXPIRIES = 365 * 1
        self.RESPONSE_LOG_FORMAT = self.RESPONSE_LOG.format
        self.SGBID_EXPIRIES = 365 * 1
        self.SGSID_EXPIRIES = None
        self.SGUUID_EXPIRIES = None

    @staticmethod
    def save_session(request):
        if request.user.is_anonymous and request.session.session_key is None:
            request.session.save()

    # to identify a session
    @staticmethod
    def set_sid(request, response):
        if request.session.get(helper_settings.logging_session_sid_name, None) is None:
            request.session[helper_settings.logging_session_sid_name] = uuid.uuid4().hex
        return request.session[helper_settings.logging_session_sid_name]

    # to identify a browser
    @staticmethod
    def set_bid(request, response):
        bid = request.COOKIES.get(helper_settings.logging_cookie_bid_name, None)
        response.set_cookie(
            helper_settings.logging_cookie_bid_name,
            domain=helper_settings.logging_cookie_domain,
            value=bid if bid is not None else uuid.uuid4().hex,
            expires=timezone.datetime.now() + timezone.timedelta(days=365)
        )
        return bid

    @staticmethod
    def set_kjzd_user_id(request, response):
        if hasattr(request, 'user'):
            if not request.user.is_anonymous:
                kjzd_user_id = request.user.kjzd_user_id
                if kjzd_user_id:
                    response.set_cookie(
                        helper_settings.logging_cookie_kjzd_user_id_name,
                        domain=helper_settings.logging_cookie_domain,
                        value=kjzd_user_id,
                        expires=None)
                    return kjzd_user_id
            else:
                response.delete_cookie(
                    helper_settings.logging_cookie_kjzd_user_id_name, domain=helper_settings.logging_cookie_domain)
        return ''

    @staticmethod
    def _get_traceback(exc_info=None):
        """Helper function to return the traceback as a string"""
        import traceback
        return '\n'.join(traceback.format_exception(*(exc_info or sys.exc_info())))

    def process_response(self, request, response):

        if request.session.get('SGUID', None) is None:
            request.session['SGUID'] = str(uuid.uuid1())

        SGUID = request.session['SGUID']
        response.set_cookie(
            'SGUID',
            value=SGUID,
            expires=self.SGUID_EXPIRIES if not self.SGUID_EXPIRIES else datetime.datetime.now() + datetime.timedelta(
                days=self.SGUID_EXPIRIES)
        )
        SGBID = request.COOKIES.get('SGBID', None)
        SGBID = SGBID if SGBID and len(SGBID) == 32 else uuid.uuid1().hex
        response.set_cookie(
            'SGBID',
            value=SGBID,
            expires=self.SGBID_EXPIRIES if not self.SGBID_EXPIRIES else datetime.datetime.now() + datetime.timedelta(
                days=self.SGBID_EXPIRIES)
        )
        if hasattr(request, 'user'):
            if not request.user.is_anonymous:
                SGUUID = request.user.kjzd_user_id
                if SGUUID:
                    response.set_cookie(
                        'SGUUID',
                        value=SGUUID,
                        expires=self.SGUUID_EXPIRIES if not self.SGUUID_EXPIRIES else datetime.datetime.now() + datetime.timedelta(
                            days=self.SGUUID_EXPIRIES)
                    )
            else:
                response.delete_cookie('SGUUID')
        SGSID = request.COOKIES.get('SGSID', None)
        if not SGSID:
            SGSID = uuid.uuid4().hex
            response.set_cookie(
                'SGSID',
                value=SGSID,
                expires=self.SGSID_EXPIRIES if not self.SGSID_EXPIRIES else datetime.datetime.now() + datetime.timedelta(
                    days=self.SGSID_EXPIRIES
                ))

        user_id = request.session.get('_auth_user_id', '')
        response.set_cookie(
            'user_id',
            value=user_id,
            expires=self.SGUUID_EXPIRIES if not self.SGUUID_EXPIRIES else datetime.datetime.now() + datetime.timedelta(
                days=self.SGUUID_EXPIRIES)
        )

        sid = self.set_sid(request, response)
        bid = self.set_bid(request, response)
        kjzd_user_id = self.set_kjzd_user_id(request, response)
        query_string = request.META.get('QUERY_STRING', None)
        log_text = self.RESPONSE_LOG_FORMAT(
            user_id=request.session.get('_auth_user_id', ''),
            ip=request.META.get('REMOTE_ADDR', ''),
            request_method=request.method,
            request_url=request.path,
            protocol=request.META.get('SERVER_PROTOCOL', ''),
            status_code=response.status_code,
            referrer=request.META.get('HTTP_REFERER', ''),
            ua=request.META.get('HTTP_USER_AGENT', ''),
            query_string='' if not query_string else ''.join(('?', query_string)),
            content_type=response.get('content-type', ''),
            sid=sid,
            bid=bid,
            kjzd_user_id=kjzd_user_id,
        )
        self.LOGGER.info(log_text)
        return response

    def process_exception(self, request, exception):
        if isinstance(exception, helper_settings.logging_except_exceptions):
            return
        try:
            request_repr = repr(request)
        except Exception as e:
            logger.warning(e)
            request_repr = "Request repr() unavailable"

        message = "{{{\n%s\n}}}\n\n{{{\n%s\n}}}" % (self._get_traceback(sys.exc_info()), request_repr)
        logger.exception(message)


class ClientAuthenticationMiddleware(object):
    def process_request(self, request):
        access_token = cache.get("CLIENT_SELLERWANT_ACCESS_TOKEN")

        if not access_token:
            try:
                data = {
                    "grant_type": "client_credentials"
                }

                reps = requests.post(settings.APP_MANAGE_URL,
                                     auth=(settings.CLIENT_ID, settings.CLIENT_SECERT),
                                     data=data,
                                     headers={"content-type": "application/x-www-form-urlencoded"})

                result = json.loads(reps.content)

                access_token = result.get("access_token")

                cache.set("CLIENT_SELLERWANT_ACCESS_TOKEN", access_token, result.get("expires_in") - 10 * 60)
            except Exception:
                pass

    def process_response(self, request, response):

        response.set_cookie('access_token', cache.get("CLIENT_SELLERWANT_ACCESS_TOKEN"))

        return response
