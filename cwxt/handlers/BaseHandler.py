import time
from tornado.web import RequestHandler
import traceback
from typing import Union, Any
from settings import TORNADO_LOCAL_SETTING


class BaseHandler(RequestHandler):

    def __init__(self, application, request, **kwargs):

        super(BaseHandler, self).__init__(application, request, **kwargs)
        self._operator_name = 'NO_LOGIN'
        self.paging = False

    @property
    def client(self):
        return self.application.pool

    @property
    def cache(self):
        return self.application.cache

    def get_current_user(self):
        pass

    def _request_summary(self):
        if self.request.method == 'post':
            return "%s %s (%s@%s)" % (self.request.method, self.request.uri,
                                      self._operator_name, self.request.remote_ip)

        return "%s %s %s(%s@%s)" % (self.request.method, self.request.uri, self.request.body.decode(),
                                    self._operator_name, self.request.remote_ip)

    def write(self, chunk: Union[str, bytes, dict, list]):
        if (isinstance(chunk, dict) and not chunk.get("code")) or isinstance(chunk, list) or isinstance(chunk, str):
            # tart
            if isinstance(chunk, list) and self.paging:
                count = len(chunk)
                page, total_page_number, _start, _end = self.get_page_re(count, 10)
                chunk = {
                    "current_page": page,
                    "total_page": total_page_number,
                    "count": count,
                    "data": chunk[_start:_end]
                }
            elif isinstance(chunk, dict) and self.paging:
                count = len(chunk['results'])
                page, total_page_number, _start, _end = self.get_page_re(count, 10)
                chunk = {
                    "current_page": page,
                    "total_page": total_page_number,
                    "count": count,
                    "data": chunk['results'][_start:_end],
                    "summary_data": chunk.get('summary_data', None)
                }
            super(BaseHandler, self).write({
                "code": 0,
                "data": chunk,
                "extra": {},
                "message": "success",
                "path": self.request.uri,
                "timestamp": self.request.request_time()
            })
        else:
            super(BaseHandler, self).write(chunk)

    def write_error(self, status_code: int, **kwargs: Any):
        if TORNADO_LOCAL_SETTING != 'prod':
            super(BaseHandler, self).write_error(status_code, **kwargs)
        else:
            message = kwargs.get("message", "error")
            error_data = {
                "code": status_code,
                "data": {},
                "extra": {},
                "message": message,
                "path": self.request.uri,
                "timestamp": self.request.request_time()
            }
            self.write(error_data)

    def get_page_re(self, count, one_page_max_number):
        """
        取得页数的方法
        :param one_page_max_number: 每页最多显示数
        :param request:
        :param count: 总数
        :return:
        """
        page = self.get_argument('page', "1")
        one_page_max_number = int(self.get_argument('count', one_page_max_number))

        total_page_number = None

        try:

            # 最大页数算出
            if count % one_page_max_number == 0:
                total_page_number = count / one_page_max_number
            else:
                total_page_number = int(count / one_page_max_number) + 1

            if total_page_number == 0:
                total_page_number = 1

            page = int(page)

            if page < 1:
                # 当页数<1的时候，设为1
                page = 1
            elif page > total_page_number:

                # 当页数>总页数的时候，设为总页数
                page = total_page_number
        except (ValueError, TypeError):

            # 当发生转换错误时，设为第1页
            page = 1
            if total_page_number is None:
                total_page_number = page
        _start = (page - 1) * int(one_page_max_number)
        _end = page * int(one_page_max_number)

        return page, total_page_number, _start, _end

