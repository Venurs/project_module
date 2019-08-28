# coding: utf-8
from tornado.web import authenticated
from tornado.gen import coroutine
from handlers.BaseHandler import BaseHandler
from utils import str2md5


class LoginHandler(BaseHandler):

    @coroutine
    def post(self):

        phone = self.get_body_argument('username')
        password = self.get_body_argument('password')
        result = yield self.client.execute("select id as user_id, phone, pwd from oa_admin WHERE phone = %s", (phone))
        user_data = [{"user_id": user_id, "phone": phone, "pwd": pwd} for user_id, phone, pwd in result.fetchall()]

        if not user_data:
            self.send_error(status_code=404, message='User Not Found!')
        if str2md5(password) == user_data[0]['pwd']:
            self.set_session('user', phone)
            self.set_secure_cookie('user', str(user_data[0]['user_id']))
            self.write('success')
        else:
            self.send_error(status_code=400, message='Wrong Password!')

    @coroutine
    def get(self):
        if self.current_user:
            self.redirect("/")
        else:
            self.render('/api/log_in/')


class IndexHandler(BaseHandler):

    @authenticated
    def get(self):
        self.write("hello %s" % self.current_user)